package dockerclienteng1

import (
	"fmt"
	"github.com/cst05001/duang/models"
	"github.com/docker/docker/api/types"
	docker "github.com/fsouza/go-dockerclient"
)

func (this *DockerClient) CreateContainer(unit *models.Unit) (types.ContainerCreateResponse, error) {
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
	containerCreateResponse := &types.ContainerCreateResponse{}
	container, err := this.Client.CreateContainer(*createContainerOptions)
	if err != nil {
		fmt.Println(err)
		containerCreateResponse.Warnings = append(containerCreateResponse.Warnings, err.Error())
		return *containerCreateResponse, err
	}
	fmt.Println(container)
	containerCreateResponse.ID = container.ID
	return *containerCreateResponse, nil
}

func (this *DockerClient) StartContainer(id string, unit *models.Unit) error {
	hostConfig := &docker.HostConfig{}
	for _, p := range unit.Parameteres {
		if p.Type == "v" {
			hostConfig.Binds = append(hostConfig.Binds, p.Value)
		}
	}
	err := this.Client.StartContainer(id, hostConfig)
	if err != nil {
		fmt.Sprintf("StartContainer: Faile with error %s\n", err)
		return err
	}
	return nil
}

func (this *DockerClient) Run(unit *models.Unit) error {
	containerCreateResponse, err := this.CreateContainer(unit)
	if err != nil {
		return err
	}
	err = this.StartContainer(containerCreateResponse.ID, unit)
	return err
}
