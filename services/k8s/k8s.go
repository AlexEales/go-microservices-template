package k8s

import (
	"context"
	"flag"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type Client struct {
	*kubernetes.Clientset
}

func MustNewClient() *Client {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.WithError(err).Fatal("error building k8s client config")
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.WithError(err).Fatal("error creating k8s client from config")
	}

	return &Client{
		Clientset: clientset,
	}
}

func (c *Client) PodsHaveStatus(ctx context.Context, selector string, status corev1.PodPhase) (bool, error) {
	pods, err := c.CoreV1().Pods("").List(ctx, v1.ListOptions{
		LabelSelector: selector,
	})
	if err != nil {
		return false, err
	}

	for _, pod := range pods.Items {
		if pod.Status.Phase != status {
			return false, nil
		}
	}
	return true, nil
}

// WaitForPodToBeReady takes a selector and a timeout string and polls the k8s cluster
// to see if the pod(s) under the specified selector are "Ready"
func (c *Client) WaitForPodToBeReady(selector, timeout string) error {
	cmd := exec.Command(
		"kubectl",
		"wait",
		"--for=condition=ready",
		fmt.Sprintf("--timeout=%s", timeout),
		"pod",
		"-l",
		selector,
	)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

// Exec performs the k8s `exec` command on the specified resource and returns
// the output of the command on success and error on fail
// Note: the command string will be split on spaces ' ' for use in exec.Command
func (c *Client) Exec(resource, command string) (string, error) {
	args := []string{
		"exec",
		resource,
		"--",
	}
	args = append(args, strings.Split(command, " ")...)
	cmd := exec.Command(
		"kubectl",
		args...,
	)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
