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
	startCmd     = exec.Command("plz", "minikube", "start")
	stopCmd      = exec.Command("plz", "minikube", "stop")
)

// Monitor is a minikube monitor
type Monitor struct{}

// NewMonitor returns a new Minikube Monitor instance to be
// used to monior the minikube cluster's health
func NewMonitor() *Monitor {
	return &Monitor{}
}

// Delete attempts to delete the minikube cluster, erroring on fail
func (m *Monitor) Delete() error {
	log.Info("Deleting Minikube")
	cmd := deleteCmd
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	log.Info("Minikube deleted successfully")
	return nil
}

// TODO: Eventually should return a status struct which contains fields mapping to plz minikube status output
// GetStatus returns the string output from `plz minikube status`
// or error if command fails to execute
func (m *Monitor) GetStatus() (string, error) {
	out, err := getStatusCmd.Output()
	return string(out), err
}

// MinikubeIsRunning returns true if minikube is running or false
// if not or an error occurs
func (m *Monitor) IsRunning() (bool, error) {
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

// Remake attempts to delete and re-start the minikube cluster,
// erroring on fail
func (m *Monitor) Remake() error {
	if err := m.Delete(); err != nil {
		return err
	}
	if err := m.Start(); err != nil {
		return err
	}
	return nil
}

// Start starts the minikube cluster if it is not running already
// erroring on failure
func (m *Monitor) Start() error {
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

// Stop attempts to stop the minikube cluster, erroring on failure
func (m *Monitor) Stop() error {
	log.Info("Stopping Minikube")
	cmd := stopCmd
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	log.Info("Minikube stopped successfully")
	return nil
}
