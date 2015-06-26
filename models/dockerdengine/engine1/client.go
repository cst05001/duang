package engine1

import (
	"fmt"
	"github.com/cst05001/duang/models/core"
	docker "github.com/fsouza/go-dockerclient"
)

type DockerClientEng1 struct {
	ClientMap map[*core.Dockerd]*docker.Client
	Unit      *core.Unit
}

func NewDockerClientEng1(unit *core.Unit) *DockerClientEng1 {
	client := &DockerClientEng1{}
	client.ClientMap = make(map[*core.Dockerd]*docker.Client)

	for _, dockerd := range unit.Dockerd {
		var err error
		c, err := docker.NewClient(dockerd.Addr)
		if err != nil {
			fmt.Printf("NewDockerClientEng1: %s\n", err)
			return nil
		}
		client.ClientMap[dockerd] = c
	}
	return client
}
