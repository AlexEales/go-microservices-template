package k8s

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

const outputFormatString = "custom-columns=Status:{.status.phase}"

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) PodIsRunning(podName string) (bool, error) {

	return false, nil
}

func (c *Client) PodsAreRunning(selector string) (bool, error) {
	out, err := exec.Command("kubectl", "get", "pods", "-l", selector, "-o", outputFormatString).Output()
	if err != nil {
		return false, err
	}

	log.Printf("Output from kubectl = %s", string(out))

	statuses := strings.Split(string(out), "\n")
	if len(statuses) <= 1 {
		return false, fmt.Errorf("no pods found for given name")
	}
	statuses = statuses[1:]

	log.Printf("Statuses = %+v", statuses)

	allRunning := true
	for _, status := range statuses {
		if strings.EqualFold(status, "Running ") {
			allRunning = false
			break
		}
	}

	return allRunning, nil
}
