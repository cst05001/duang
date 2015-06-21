package dockerclienteng1

import (
	"fmt"
	docker "github.com/fsouza/go-dockerclient"
)

type DockerClient struct {
	Client *docker.Client
}

func NewDockerClient(endpoint string) *DockerClient {
	client := &DockerClient{}
	var err error
	client.Client, err = docker.NewClient(endpoint)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return client
}
