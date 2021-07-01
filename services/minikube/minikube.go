package minikube

import (
	"errors"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

var (
	getStatusCmd = exec.Command("plz", "minikube", "status")
	startCmd     = exec.Command("plz", "minikube", "start")
)

// Monitor is the interface for a Minikube Monitor
type Monitor interface {
	GetStatus() (string, error)
	IsRunning() (bool, error)
	Start() error
}

type monitor struct{}

// NewMonitor returns a new Minikube Monitor instance to be
// used to monior the minikube cluster's health
func NewMonitor() Monitor {
	return &monitor{}
}

// TODO: Eventually should return a status struct which contains fields mapping to plz minikube status output
// GetStatus returns the string output from `plz minikube status`
// or error if command fails to execute
func (m *monitor) GetStatus() (string, error) {
	out, err := getStatusCmd.Output()
	return string(out), err
}

// MinikubeIsRunning returns true if minikube is running or false
// if not or an error occurs
func (m *monitor) IsRunning() (bool, error) {
	_, err := m.GetStatus()

	if err == nil {
		return true, nil
	}

	exitErr := &exec.ExitError{}
	if !errors.As(err, &exitErr) || exitErr.ExitCode() != 85 {
		return false, err
	}

	return false, nil
}

// Start starts the minikube cluster if it is not running already
// erroring on failure
func (m *monitor) Start() error {
	isRunning, err := m.IsRunning()
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
