package dockerclienteng1

import (
	"fmt"
	"github.com/cst05001/duang/models"
	docker "github.com/fsouza/go-dockerclient"
)

func (this *DockerClient) CreateContainer(unit *models.Unit) {
	hostConfig := &docker.HostConfig{}
	for _, p := range unit.Parameteres {
		if p.Type == "v" {
			hostConfig.Binds = append(hostConfig.Binds, p.Value)
		}
	}
	config := &docker.Config{
		Image: unit.Image,
	}
	createContainerOptions := &docker.CreateContainerOptions{
		Name:       unit.Name,
		Config:     config,
		HostConfig: hostConfig,
	}
	container, err := this.Client.CreateContainer(*createContainerOptions)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(container)
}
