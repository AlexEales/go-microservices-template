# Go Microservices Template

Template repository designed to let you get up and running with a go microservices architecture quickly and easily.

## Contents

- [Go Microservices Template](#go-microservices-template)
  - [Contents](#contents)
  - [Prerequisits](#prerequisits)
    - [Suggested Tools](#suggested-tools)
  - [Features](#features)
  - [Usage](#usage)
    - [Project Configuration and Initial Setup](#project-configuration-and-initial-setup)
    - [Developing using this Template](#developing-using-this-template)
    - [Deploying using this Template](#deploying-using-this-template)

## Prerequisits

It is recommended to familiarize yourself with the following tools before getting up and running with this template as the build system and structure might not be conventional:

- [Golang](https://golang.org/) (not required as can be installed by project but useful)
- [Plz Build Tool](https://please.build/) (again not required can be installed using the `pleasew` file in repo)
  - Recommend doing the codelabs on the website if you haven't used Please before.
- [Docker](https://www.docker.com/) Required on host machine
- [Kubernetes](https://kubernetes.io/) Required on host machine

### Suggested Tools

These are some of the tools I make use of when developing microservices with go which might be useful have a look for yourself and decide.

- [K9s](https://github.com/derailed/k9s) - Upgraded command line K8s replacement with handy port forwarding and debugging tools for your kubernetes cluster(s).

## Features

This project is still WIP and so not all features will be present yet but the ones completed or being developed will be marked appropriately:

- [ ] Common libraries for DB access etc.
- [x] Common 3rd Party Dependencies for Building and Testing Microservices (see `third_party/go/BUILD`)
- [ ] Tool for updating `go.mod` file with dependencies in `third_party/go/BUILD` to help with autocompletion in VSCode and other IDEs
- [ ] Postgres DB Deployment(s)
  - [ ] SQL Migration K8s Job/Init Container
- [ ] Kafka Deployment
- [ ] Hashicorp Vault Deployment
- [ ] Prometheus, Kibana, OpenTracing and Grafana Observability Stack Deployments
- [ ] gRPC Gateway Buildrules
- [ ] Example Services
  - [ ] HTTP
  - [ ] gRPC
  - [ ] Kafka
  - All make use of the DB and Hashicorp Vault
- [ ] Example Product Deployment Script
  - Configurable deployment script for deploying services and your created binaries in order and verifying the deployments.

> All **Deployments** are not required and the project only contains the k8s files/builds for deploying these to your local cluster you are free to delete these if you do not require them. This will howerver also require potential deletions in the `third_party/go/BUILD` file as some dependencies might linger which will no longer be used but this will not cause your deployments or builds to be larger than if they were removed.

## Usage

After templating your new repository, follow these steps to get it setup for use and verify everything is working as expected:

### Project Configuration and Initial Setup

1. Edit `.plzconfig` and `go.mod` and replace any occurances of `go-microservices-template` with your project's name.
2. If you wish to pin the versions of the included build definitions in `build/defs` enter the folder and, as the comment prescribes, swap the `master` part of the URL for the commit hash you with to pin.
3. Run `plz build //tools/...` which will build your own isolated go toolchain and download the latest version of minikube.
4. Run `plz minikube start` to get a local instance of minikube up and running and once completed run `eval $(plz minikube docker-env)` to link our docker env to minikube (hopefully will be automated in the future).

### Developing using this Template

### Deploying using this Template
