package engine1

import (
	"fmt"
	"github.com/cst05001/duang/models/core"
	docker "github.com/fsouza/go-dockerclient"
)

type DockerClientEng1 struct {
	Client []*docker.Client
	Unit   *core.Unit
}

func NewDockerClientEng1(unit *core.Unit) *DockerClientEng1 {
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
