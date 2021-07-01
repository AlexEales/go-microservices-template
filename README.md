# Go Microservices Template

Template repository designed to let you get up and running with a go microservices architecture quickly and easily.

## Contents

- [Go Microservices Template](#go-microservices-template)
  - [Contents](#contents)
  - [Prerequisits](#prerequisits)
    - [Suggested Tools](#suggested-tools)
  - [Features](#features)
  - [Usage](#usage)

## Prerequisits

It is recommended to familiarize yourself with the following tools before getting up and running with this template as the build system and structure might not be conventional:

- [Golang](https://golang.org/) (not required as can be installed by project but useful)
- [Plz Build Tool](https://please.build/) (again not required can be installed using the `pleasew` file in repo)
- [Docker](https://www.docker.com/) Required on host machine
- [Kubernetes](https://kubernetes.io/) Required on host machine

### Suggested Tools

These are some of the tools I make use of when developing microservices with go which might be useful have a look for yourself and decide.

- [K9s](https://github.com/derailed/k9s) - Upgraded command line K8s replacement with handy port forwarding and debugging tools for your kubernetes cluster(s).

## Features

This project is still WIP and so not all features will be present yet but the ones completed or being developed will be marked appropriately:

- [ ] Common libraries for DB access etc.
- [ ] Postgres DB Deployment(s)
- [ ] Kafka Deployment
- [ ] Hashicorp Vault Deployment
- [ ] Prometheus and Grafana Observability Stack Deployments
- [ ] OpenTracing Deployment
- [ ] Kibana Deployment
- [ ] Example Services
  - [ ] HTTP
  - [ ] gRPC
  - [ ] Kafka
  - All make use of the DB and Hashicorp Vault

> All **Deployments** are not required and the project only contains the k8s files/builds for deploying these to your local cluster you are free to delete these if you do not require them.

## Usage

After templating your new repository, follow these steps to get it setup for use and verify everything is working as expected:

1. Edit `.plzconfig` and `go.mod` and replace any occurances of `go-microservices-template` with your project's name.
2. If you wish to pin the versions of the included build definitions in `build/defs` enter the folder and, as the comment prescribes, swap the `master` part of the URL for the commit hash you with to pin.
3. Run `plz build //tools/...` which will build your own isolated go toolchain and download the latest version of minikube.
4. Run `plz minikube start` to get a local instance of minikube up and running and once completed run `eval $(plz minikube docker-env)` to link our docker env to minikube (hopefully will be automated in the future).
