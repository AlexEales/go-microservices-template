package main

import (
	"go-microservices-template/services/helm"
	"go-microservices-template/services/minikube"

	log "github.com/sirupsen/logrus"
)

// TODO: Add in configurability for the configs etc. and add in retries to wait for pods etc. to be deployed

const (
	consulValuesConfigPath = "services/vault/helm-consul-values.yaml"
	vaultValuesConfigPath  = "services/vault/helm-vault-values.yaml"
)

func main() {
	// TODO: Eventually this will be a common cluster monitor to deploy to any cluster not just minikube
	minikubeMonitor := minikube.NewMonitor()
	if err := minikubeMonitor.Start(); err != nil {
		log.WithError(err).Fatal("error checking or starting minikube")
	}

	helmClient := helm.NewClient()
	if installed, err := helmClient.ChartsInstalled("consul", "vault"); err != nil {
		log.WithError(err).Fatal("error determining if helm charts are already installed")
	} else if installed {
		log.Info("Helm charts already installed, removed consul and vault helm charts and re-run")
		return
	}

	if err := helmClient.AddRepository("hashicorp", "https://helm.releases.hashicorp.com"); err != nil {
		log.WithError(err).Fatal("error adding Hashicorp helm repository")
	}

	if err := helmClient.UpdateRepositories(); err != nil {
		log.WithError(err).Fatal("error updating Hashicorp helm repository")
	}

	// TODO: Need k8s utils for checking if these charts have installed correctly with backoff

	if err := helmClient.InstallChart("hashicorp", "consul", "--values", consulValuesConfigPath); err != nil {
		log.WithError(err).Fatal("error installing Hashicorp Consul helm chart")
	}

	if err := helmClient.InstallChart("hashicorp", "vault", "--values", vaultValuesConfigPath); err != nil {
		log.WithError(err).Fatal("error installing Hashicorp Consul helm chart")
	}
}
