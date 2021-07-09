package main

import (
	"fmt"
	"go-microservices-template/common/go/helm"
	"go-microservices-template/common/go/minikube"
	"go-microservices-template/common/go/retry"

	log "github.com/sirupsen/logrus"
)

const (
	openingASCIIArt = `
    ____           __        _____            
   /  _/___  _____/ /_____ _/ / (_)___  ____ _
   / // __ \/ ___/ __/ __ '/ / / / __ \/ __ '/
 _/ // / / (__  ) /_/ /_/ / / / / / / / /_/ / 
/___/_/ /_/____/\__/\__,_/_/_/_/_/ /_/\__, /  
 _    __            ____             /____/   
| |  / /___ ___  __/ / /_                     
| | / / __ '/ / / / / __/                     
| |/ / /_/ / /_/ / / /_                       
|___/\__,_/\__,_/_/\__/                       
`
	closingASCIIArt = `
 _    __            ____                    
| |  / /___ ___  __/ / /_                   
| | / / __ '/ / / / / __/                   
| |/ / /_/ / /_/ / / /_                     
|___/\__,_/\__,_/_/\__/      ____         __
   /  _/___  _____/ /_____ _/ / /__  ____/ /
   / // __ \/ ___/ __/ __ '/ / / _ \/ __  / 
 _/ // / / (__  ) /_/ /_/ / / /  __/ /_/ /  
/___/_/ /_/____/\__/\__,_/_/_/\___/\__,_/   
`
)

func startMinikubeIfNotRunning() error {
	client := minikube.NewClient()

	log.Info("Checking if minikube is running...")
	running, err := client.IsRunning()
	if err != nil {
		return err
	}

	if running {
		log.Info("Minikube already running")
		return nil
	}

	log.Info("Minikube not running, starting minikube...")
	if err := client.Start(); err != nil {
		return err
	}
	return nil
}

func initialiseHelm(client helm.Client) error {
	if err := client.AddRepository("hashicorp", "https://helm.releases.hashicorp.com"); err != nil {
		return err
	}

	if err := client.UpdateRepositories(); err != nil {
		return err
	}

	return nil
}

func waitForHelmChartToInstall(isInstalled func() (bool, error), retryOpts *retry.Opts) error {
	return nil
}

func main() {
	fmt.Println(openingASCIIArt)

	if err := startMinikubeIfNotRunning(); err != nil {
		log.WithError(err).
			Fatal("error checking if minikube is running or starting minikube")
	}

	helmClient := helm.NewClient()
	installedMatrix, err := helmClient.ChartsInstalled("consul", "vault")
	if err != nil {
		log.WithError(err).
			Fatal("error checking if Consul and Vault Helm charts are installed")
	}

	if installedMatrix == nil {
		log.Info("Initialising Helm and adding Hashicorp Helm repository")
		if err := initialiseHelm(helmClient); err != nil {
			log.WithError(err).
				Fatal("error initialising helm")
		}
	} else {
		log.Info("Hashicorp Helm repository already installed, skipping helm initialisation")
	}

	if !installedMatrix["consul"] {
		log.Info("Installing Consul Helm chart...")
	}

	if !installedMatrix["vault"] {
		log.Info("Installing Vault Helm chart...")
	}

	fmt.Println(closingASCIIArt)
}
