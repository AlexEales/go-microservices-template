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
type Client struct{}

// NewClient returns a new helm client
func NewClient() *Client {
	return &Client{}
}

// AddRepository adds a helm repository
func (c *Client) AddRepository(repository, url string) error {
	log.Infof("Installing %s helm repository", strings.Title(repository))
	cmd := exec.Command("plz", "helm", "repo", "add", repository, url)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	log.Infof("%s helm repository installed successfully", strings.Title(repository))
	return nil
}

// InstallChart installs a new helm chart on the current k8s context
func (c *Client) InstallChart(repository, chart string, args ...string) error {
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
	log.Infof("%s %s helm chart installed successfully", strings.Title(repository), strings.Title(chart))
	return nil
}

// UpdateRepositories updates the added helm repositories
func (c *Client) UpdateRepositories() error {
	log.Info("Updating helm repositories")
	cmd := updateRepositoriesCmd
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	log.Info("Helm repositories updated")
	return nil
}
