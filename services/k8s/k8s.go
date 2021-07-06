package k8s

import (
	"fmt"
	"os/exec"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
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
func (c *Client) Exec(resource, command string) (string, error) {
	cmd := exec.Command(
		"kubectl",
		"exec",
		resource,
		"--",
		command,
	)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
