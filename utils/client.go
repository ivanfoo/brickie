package utils

import (
	"log"

	"github.com/fsouza/go-dockerclient"
)

func NewDockerClient() *docker.Client {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	return client
}
