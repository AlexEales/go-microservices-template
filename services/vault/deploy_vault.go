package main

import (
	"context"
	"fmt"
	"go-microservices-template/common/go/backoff"
	"go-microservices-template/common/go/retry"
	"go-microservices-template/services/helm"
	"go-microservices-template/services/k8s"
	"go-microservices-template/services/minikube"
	"time"

	flags "github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

var (
	errPodNotRunning = fmt.Errorf("pod not running")

	k8sBackoffOpts = &backoff.Exponential{
		InitialDelay: 4 * time.Second,
		BaseDelay:    4 * time.Second,
		MaxDelay:     64 * time.Second,
		Factor:       2,
		Jitter:       0,
	}
	k8sRetryOpts = &retry.Opts{
		MaxAttempts: 5,
		Backoff:     k8sBackoffOpts,
		IsRetryable: func(e error) bool {
			return e == errPodNotRunning
		},
	}

	consulK8sSelector = "app=consul"
	vaultK8sSelector  = "app.kubernetes.io/name=vault"
)

var opts struct {
	ConsulConfigPath string `long:"consul-config-path" default:"services/vault/config/helm-consul-values.yaml" env:"CONSUL_CONFIG_PATH"`
	VaultConfigPath  string `long:"vault-config-path" default:"services/vault/config/helm-vault-values.yaml" env:"VAULT_CONFIG_PATH"`
}

func podRunningRetryFn(selector string) func() error {
	return func() error { return nil }
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		log.WithError(err).Fatal("failed to parse command line configuration")
	}

	// TODO: Check the config files exist?

	// TODO: Eventually this will be a common cluster monitor to deploy to any cluster not just minikube
	minikubeMonitor := minikube.NewMonitor()
	if err := minikubeMonitor.Start(); err != nil {
		log.WithError(err).Fatal("error checking or starting minikube")
	}

	helmClient := helm.NewClient()
	// TODO: maybe change installed to return a map from chart to bool for later skipping?
	if installed, err := helmClient.ChartsInstalled("consul", "vault"); err != nil {
		log.WithError(err).Fatal("error determining if helm charts are already installed")
	} else if installed {
		// TODO: not sure if this is entirely what we would want to do as we might want to check the pods have been initialised etc.
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

	if err := helmClient.InstallChart("hashicorp", "consul", "--values", opts.ConsulConfigPath); err != nil {
		log.WithError(err).Fatal("error installing Hashicorp Consul helm chart")
	}

	k8sClient := k8s.NewClient()
	checkConsulRunning := func() error {
		isRunning, err := k8sClient.PodsAreRunning(consulK8sSelector)
		if err != nil {
			return err
		}
		if isRunning {
			return nil
		}
		return errPodNotRunning
	}
	if _, err := retry.Do(context.Background(), checkConsulRunning, k8sRetryOpts); err != nil {
		log.WithError(err).Fatal("error occured waiting for consul to be deployed")
	}

	if err := helmClient.InstallChart("hashicorp", "vault", "--values", opts.VaultConfigPath); err != nil {
		log.WithError(err).Fatal("error installing Hashicorp Consul helm chart")
	}

	checkVaultRunning := func() error {
		isRunning, err := k8sClient.PodsAreRunning(vaultK8sSelector)
		if err != nil {
			return err
		}
		if isRunning {
			return nil
		}
		return errPodNotRunning
	}
	if _, err := retry.Do(context.Background(), checkVaultRunning, k8sRetryOpts); err != nil {
		log.WithError(err).Fatal("error occured waiting for vault to be deployed")
	}
}
