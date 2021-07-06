package main

import (
	"context"
	"errors"
	"os/exec"
	"time"

	flags "github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"

	"go-microservices-template/common/go/backoff"
	"go-microservices-template/common/go/retry"
	"go-microservices-template/services/helm"
	"go-microservices-template/services/k8s"
	"go-microservices-template/services/minikube"
)

const (
	consulK8sSelector             = "app=consul"
	vaultAgentInjectorK8sSelector = "app.kubernetes.io/name=vault-agent-injector"
)

var (
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
		IsRetryable: func(err error) bool {
			if err == nil {
				return false
			}
			var exitErr *exec.ExitError
			if errors.As(err, &exitErr) {
				if exitErr.ExitCode() == 1 {
					return true
				}
			}
			return false
		},
	}
)

var opts struct {
	ConsulConfigPath string `long:"consul-config-path" default:"services/vault/config/helm-consul-values.yaml" env:"CONSUL_CONFIG_PATH"`
	VaultConfigPath  string `long:"vault-config-path" default:"services/vault/config/helm-vault-values.yaml" env:"VAULT_CONFIG_PATH"`
}

func initialiseHelm(client *helm.Client) {
	// TODO: maybe change installed to return a map from chart to bool for later skipping?
	if installed, err := client.ChartsInstalled("consul", "vault"); err != nil {
		log.WithError(err).Fatal("error determining if helm charts are already installed")
	} else if installed {
		// TODO: not sure if this is entirely what we would want to do as we might want to check the pods have been initialised etc.
		log.Info("Helm charts already installed, re-run on fresh minikube instance")
		return
	}

	if err := client.AddRepository("hashicorp", "https://helm.releases.hashicorp.com"); err != nil {
		log.WithError(err).Fatal("error adding Hashicorp helm repository")
	}

	if err := client.UpdateRepositories(); err != nil {
		log.WithError(err).Fatal("error updating Hashicorp helm repository")
	}
}

func waitForConsulToInstall(helmClient *helm.Client, k8sClient *k8s.Client) {
	if err := helmClient.InstallChart("hashicorp", "consul", "--values", opts.ConsulConfigPath); err != nil {
		log.WithError(err).Fatal("error installing Hashicorp Consul helm chart")
	}

	log.Info("Waiting Consul pods to be ready")
	checkConsulRunning := func() error {
		return k8sClient.WaitForPodToBeReady(consulK8sSelector, "15s")
	}
	if _, err := retry.Do(context.Background(), checkConsulRunning, k8sRetryOpts); err != nil {
		log.WithError(err).Fatal("error occured waiting for consul to be deployed")
	}
}

func waitForVaultToInstall(helmClient *helm.Client, k8sClient *k8s.Client) {
	if err := helmClient.InstallChart("hashicorp", "vault", "--values", opts.VaultConfigPath); err != nil {
		log.WithError(err).Fatal("error installing Hashicorp Consul helm chart")
	}

	log.Info("Waiting Vault Agent Injector pod to be ready")
	checkVaultAgentInjectorRunning := func() error {
		return k8sClient.WaitForPodToBeReady(vaultAgentInjectorK8sSelector, "15s")
	}
	if _, err := retry.Do(context.Background(), checkVaultAgentInjectorRunning, k8sRetryOpts); err != nil {
		log.WithError(err).Fatal("error occured waiting for the vault agent injector to be deployed")
	}
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
	initialiseHelm(helmClient)

	k8sClient := k8s.NewClient()
	waitForConsulToInstall(helmClient, k8sClient)
	waitForVaultToInstall(helmClient, k8sClient)

	// TODO: Write k8s client method for getting all pod names for a selector so we can loop over them here
	initOutput, err := k8sClient.Exec("vault-0", "vault operator init -key-shares=1 -key-threshold=1 -format=json")
	if err != nil {
		log.WithError(err).Fatal("error initialising vault-0")
	}
	log.Infof("vault-0 initialised:\n%s", initOutput)
}
