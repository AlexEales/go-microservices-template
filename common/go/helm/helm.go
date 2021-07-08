package helm

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	updateRepositoriesCmd = exec.Command("plz", "helm", "repo", "update")
)

// Client is a helm client for interacting with helm
type Client interface {
	AddRepository(repository, url string) error
	ChartInstalled(chart string) (bool, error)
	ChartsInstalled(charts ...string) (map[string]bool, error)
	InstallChart(repository, chart string, args ...string) error
	UpdateRepositories() error
}

type client struct{}

// NewClient returns a new helm client
func NewClient() Client {
	return &client{}
}

// AddRepository adds a helm repository
func (c *client) AddRepository(repository, url string) error {
	log.Infof("Installing %s helm repository", strings.Title(repository))
	cmd := exec.Command("plz", "helm", "repo", "add", repository, url)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	log.Infof("%s helm repository installed successfully", strings.Title(repository))
	return nil
}

// ChartInstalled checks if a specified chart has been installed already to prevent installing twice
// and causing an error
func (c *client) ChartInstalled(chart string) (bool, error) {
	cmd := exec.Command("plz", "helm", "list", "-q")
	out, err := cmd.Output()
	if err != nil {
		return false, err
	}

	charts := strings.Split(string(out), "\n")
	for _, existingChart := range charts {
		if strings.EqualFold(chart, existingChart) {
			return true, nil
		}
	}
	return false, nil
}

// ChartsInstalled checks if specified charts are installed already to prevent installing twice
// and causing an error
func (c *client) ChartsInstalled(charts ...string) (map[string]bool, error) {
	var err error
	installedMatrix := make(map[string]bool, len(charts))
	for _, chart := range charts {
		installedMatrix[chart], err = c.ChartInstalled(chart)
		if err != nil {
			return nil, err
		}
	}
	return installedMatrix, nil
}

// InstallChart installs a new helm chart on the current k8s context
func (c *client) InstallChart(repository, chart string, args ...string) error {
	log.Infof("Installing %s %s helm chart", strings.Title(repository), strings.Title(chart))
	cmdArgs := []string{
		"helm",
		"install",
		chart,
		fmt.Sprintf("%s/%s", repository, chart),
	}
	cmdArgs = append(cmdArgs, args...)
	cmd := exec.Command("plz", cmdArgs...)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	log.Infof("%s %s helm chart installed", strings.Title(repository), strings.Title(chart))
	return nil
}

// UpdateRepositories updates the added helm repositories
func (c *client) UpdateRepositories() error {
	log.Info("Updating helm repositories")
	cmd := updateRepositoriesCmd
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	log.Info("Helm repositories updated")
	return nil
}
