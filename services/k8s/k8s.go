package k8s

import (
	"fmt"
	"os/exec"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

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

func (c *Client) Exec(resource, command string) (string, error) {
	return "", nil
}
