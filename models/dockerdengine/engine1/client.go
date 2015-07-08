package engine1

import (
	"github.com/astaxie/beego"
	"github.com/cst05001/duang/models/core"
	docker "github.com/fsouza/go-dockerclient"
)

type DockerClientEng1 struct {
	ClientMap map[*core.Dockerd]*docker.Client
	Unit      *core.Unit
}

func (this *DockerClientEng1) newClient(addr string) *docker.Client {
	client, err := docker.NewClient(addr)
	if err != nil {
		beego.Error(err)
		return nil
	}
	return client
}

func NewDockerClient() *DockerClientEng1 {
	client := &DockerClientEng1{}
	/*
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
	*/
	return client
}
