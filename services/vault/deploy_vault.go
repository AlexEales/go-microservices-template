package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

// TODO: Add in configurability for the configs etc. and add in retries to wait for pods etc. to be deployed

const (
	consulValuesConfigPath = "services/vault/helm-consul-values.yaml"
	vaultValuesConfigPath  = "services/vault/helm-vault-values.yaml"
)

var log = logrus.New()

// TODO: Move these minikube utilities functions into a common/services library

func getMinikubeStatus() (string, error) {
	cmd := exec.Command("plz", "minikube", "status")
	out, err := cmd.Output()
	return string(out), err
}

func minikubeIsRunning() (bool, error) {
	_, err := getMinikubeStatus()

	if err == nil {
		return true, nil
	}

	exitErr := &exec.ExitError{}
	if !errors.As(err, &exitErr) || exitErr.ExitCode() != 85 {
		return false, err
	}

	return false, nil
}

func startMinikubeIfNotRunning() error {
	isRunning, err := minikubeIsRunning()
	if err != nil {
		return err
	}

	if !isRunning {
		log.Info("Minikube not running, starting minikube")
		cmd := exec.Command("plz", "minikube", "start")
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

// TODO: Move and commonise these helm utility functions into a common/service library

func addHashicorpHelmRepo() error {
	log.Info("Adding Hashicorp helm repository")
	cmd := exec.Command("plz", "helm", "repo", "add", "hashicorp", "https://helm.releases.hashicorp.com")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	log.Info("Hashicorp helm repository added successfully")
	return nil
}

func updateHelmRepositories() error {
	log.Info("Updating helm repositories")
	cmd := exec.Command("plz", "helm", "repo", "update")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	log.Info("Helm repositories updated")
	return nil
}

func installHelmChart(repo, chart string, args ...string) error {
	log.Infof("Installing %s %s helm chart", strings.Title(repo), strings.Title(chart))
	cmdArgs := []string{
		"helm",
		"install",
		chart,
		fmt.Sprintf("%s/%s", repo, chart),
	}
	cmdArgs = append(cmdArgs, args...)
	cmd := exec.Command("plz", cmdArgs...)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	log.Infof("%s %s helm chart installed successfully", strings.Title(repo), strings.Title(chart))
	return nil
}

func main() {
	if err := startMinikubeIfNotRunning(); err != nil {
		log.WithError(err).Fatal("error checking or starting minikube")
	}

	if err := addHashicorpHelmRepo(); err != nil {
		log.WithError(err).Fatal("error adding Hashicorp helm repository")
	}

	if err := updateHelmRepositories(); err != nil {
		log.WithError(err).Fatal("error updating Hashicorp helm repository")
	}

	if err := installHelmChart("hashicorp", "consul", "--values", consulValuesConfigPath); err != nil {
		log.WithError(err).Fatal("error installing Hashicorp Consul helm chart")
	}

	if err := installHelmChart("hashicorp", "vault", "--values", vaultValuesConfigPath); err != nil {
		log.WithError(err).Fatal("error installing Hashicorp Consul helm chart")
	}
}
