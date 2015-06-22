package dockerclienteng1

import (
	"fmt"
	"github.com/cst05001/duang/models"
	docker "github.com/fsouza/go-dockerclient"
)

type DockerClientEng1 struct {
	Client []*docker.Client
	Unit   *models.Unit
}

func NewDockerClientEng1(unit *models.Unit) *DockerClientEng1 {
	client := &DockerClientEng1{}
	client.Client = make([]*docker.Client, 0)

	for _, dockerd := range unit.Dockerd {
		var err error
		c, err := docker.NewClient(dockerd.Addr)
		if err != nil {
			fmt.Printf("NewDockerClientEng1: %s\n", err)
			return nil
		}
		client.Client = append(client.Client, c)
	}
	return client
}
