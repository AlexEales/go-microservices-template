package db

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"time"

	"github.com/jessevdk/go-flags"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	"go-microservices-template/common/go/backoff"
	"go-microservices-template/common/go/retry"
)

var DefaultRetryOpts = &retry.Opts{
	MaxAttempts: 5,
	IsRetryable: func(e error) bool {
		return e != nil
	},
	Backoff: &backoff.Exponential{
		InitialDelay: 2 * time.Second,
		BaseDelay:    2 * time.Second,
		MaxDelay:     64 * time.Second,
		Factor:       2,
		Jitter:       0,
	},
}

// Tx is a wrapper around the sqlx.Tx type to allow consumers of this library to not have to import sqlx
type Tx = sqlx.Tx

// TxFn represents a function to be run in a transaction
type TxFn = func(*Tx) (sql.Result, error)

// RawConn is a wrapper around sqlx.DB
type RawConn = sqlx.DB

// DB is the interface defining the common database (derived from sqlx but could be based on any DB library)
type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	InTxn(ctx context.Context, fn TxFn) (sql.Result, error)
	MustExec(query string, args ...interface{}) sql.Result
	MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result
	Ping() error
	PingContext(ctx context.Context) error
	Select(dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

// CommonDB is a wrapper around the RawConn which implements some of the utility methods to be a DB instance
type CommonDB struct {
	*RawConn
}

// InTxn takes a function and runs the it in a transaction, handling commiting and rolling-back the result
func (db *CommonDB) InTxn(ctx context.Context, fn TxFn) (sql.Result, error) {
	tx, err := db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
		ReadOnly:  false,
	})
	if err != nil {
		return nil, err
	}

	result, err := fn(tx)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}
	return result, nil
}

// NewCommonDB takes a raw connection and returns a CommonDB instance
func NewCommonDB(rawConn *RawConn) *CommonDB {
	return &CommonDB{
		RawConn: rawConn,
	}
}

// DBOpts represents the options that can be provided to connect to a database
type DBOpts struct {
	Name string `short:"n" description:"The name of the database" required:"true" env:"DB_NAME"`
	Host string `short:"h" description:"The address of the host the DB is running on" required:"true" env:"DB_HOST"`
	Port uint32 `short:"p" description:"The port the DB is running on" default:"5432" env:"DB_PORT"`
	User string `short:"u" description:"The user connection to the database" required:"true" env:"DB_USER"`
	// TODO: replace with ssl?
	Password string `long:"password" description:"The password for connection to the database" required:"true" env:"DB_PASSWORD"`
}

// Validate checks that the opts are populated as required for connecting to a db
func (o *DBOpts) Validate() error {
	if o.Name == "" {
		return fmt.Errorf("DB name cannot be empty string")
	}

	if _, err := url.Parse(o.Host); err != nil {
		return err
	}

	if o.User == "" {
		return fmt.Errorf("user cannot be empty string")
	}

	return nil
}

// FormConnectionString takes the opts and forms a database connection string from their values
func (o *DBOpts) FormConnectionString() string {
	return fmt.Sprintf(
		"database=%s host=%s port=%d user=%s password=%s sslmode=disable",
		o.Name,
		o.Host,
		o.Port,
		o.User,
		o.Password,
	)
}

// GetDBOpts returns a DBOpts instance populated from values passed in command line args and the environment
func GetDBOpts() (*DBOpts, error) {
	opts := &DBOpts{}
	_, err := flags.Parse(opts)
	if err != nil {
		return nil, err
	}
	return opts, nil
}

// Conn represents a database connection
type Conn struct {
	DB          DB
	IsRetryable map[pq.ErrorCode]bool
	RetryOpts   *retry.Opts
}

// Connect takes database options and attempts to connect to the DB, returning an error on failue
func Connect(ctx context.Context, opts *DBOpts) (*Conn, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	var db *sqlx.DB
	var err error
	connStr := opts.FormConnectionString()
	_, err = retry.Do(ctx, func() error {
		db, err = sqlx.ConnectContext(ctx, "postgres", connStr)
		if err != nil {
			return err
		}
		return nil
	}, DefaultRetryOpts)
	if err != nil {
		return nil, err
	}

	conn := &Conn{
		DB:          NewCommonDB(db),
		IsRetryable: map[pq.ErrorCode]bool{},
		RetryOpts:   DefaultRetryOpts,
	}
	return conn, nil
}

// MustConnect takes database options and attempts to connect to the DB, panicing on error
func MustConnect(ctx context.Context, opts *DBOpts) *Conn {
	conn, err := Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	return conn
}
