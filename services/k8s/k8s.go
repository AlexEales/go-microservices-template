package k8s

import (
	"fmt"
	"log"
	"os/exec"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) WaitForResourceToBeReady(selector, timeout string) error {
	cmd := exec.Command(
		"kubectl",
		"wait",
		"--for=condition=ready",
		fmt.Sprintf("--timeout=%s", timeout),
		"pod",
		"-l",
		selector,
	)
	if _, err := cmd.Output(); err != nil {
		return err
	}
	return nil
}

func (c *Client) WaitForCondition(selector, condition, timeoutStr string) error {
	cmd := exec.Command(
		"kubectl",
		"wait",
		fmt.Sprintf("--for=condition=%s", condition),
		fmt.Sprintf("--timeout=%s", timeoutStr),
		selector,
	)
	if out, err := cmd.Output(); err != nil {
		log.Printf(string(out))
		return err
	}
	return nil
}

func (c *Client) Exec(resource, command string) (string, error) {
	return "", nil
}
