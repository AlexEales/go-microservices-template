package main

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"go-microservices-template/services/k8s"
)

func waitForPodsStatus(phase corev1.PodPhase) {}

func main() {
	client := k8s.MustNewClient()

	podsClient := client.CoreV1().Pods("default")

	for {
		pods, err := podsClient.List(context.TODO(), v1.ListOptions{
			LabelSelector: "app.kubernetes.io/name=vault",
		})
		if err != nil {
			panic(err.Error())
		}

		for _, pod := range pods.Items {
			log.WithField("phase", pod.Status.Phase).Infof("%s's status phase", pod.Name)
		}

		time.Sleep(5 * time.Second)
	}
}
