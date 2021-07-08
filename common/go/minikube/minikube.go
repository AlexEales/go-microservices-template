package minikube

import (
	"errors"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

var (
	deleteCmd    = exec.Command("plz", "minikube", "delete")
	getStatusCmd = exec.Command("plz", "minikube", "status")
	startCmd     = exec.Command("plz", "minikube", "start", "--addons", "registry", "metrics-server", "dashboard")
	stopCmd      = exec.Command("plz", "minikube", "stop")
)

// Manager defines the interface of a object which can manage a minikube cluster
type Manager interface {
	Delete() error
	Remake() error
	Start() error
	Stop() error
}

// Monitor defines the interface of a object which can monitor a minikube cluster
type Monitor interface {
	IsRunning() (bool, error)
}

// Client defines a interface of a object which can monitor and manage a minikube cluster
type Client interface {
	Manager
	Monitor
}

// NewClient returns a new minikube Client instance
func NewClient() Client {
	return &client{}
}

type client struct{}

// Delete attempts to delete the minikube cluster, erroring on fail
func (c *client) Delete() error {
	log.Info("Deleting Minikube")
	cmd := deleteCmd
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	log.Info("Minikube deleted successfully")
	return nil
}

// MinikubeIsRunning returns true if minikube is running or false
// if not or an error occurs
func (c *client) IsRunning() (bool, error) {
	err := getStatusCmd.Run()
	if err == nil {
		return true, nil
	}

	exitErr := &exec.ExitError{}
	if !errors.As(err, &exitErr) || exitErr.ExitCode() != 85 {
		return false, err
	}

	return false, nil
}

// Remake attempts to delete and re-start the minikube cluster,
// erroring on fail
func (c *client) Remake() error {
	if err := c.Delete(); err != nil {
		return err
	}
	if err := c.Start(); err != nil {
		return err
	}
	return nil
}

// Start starts the minikube cluster if it is not running already
// erroring on failure
func (c *client) Start() error {
	isRunning, err := c.IsRunning()
	if err != nil {
		return err
	}

	if !isRunning {
		log.Info("Minikube not running, starting minikube")
		cmd := startCmd
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			return err
		}
		log.Info("Minikube started")
	} else {
		log.Info("Minikube already running")
	}

	return nil
}

// Stop attempts to stop the minikube cluster, erroring on failure
func (c *client) Stop() error {
	log.Info("Stopping Minikube")
	cmd := stopCmd
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	log.Info("Minikube stopped successfully")
	return nil
}
