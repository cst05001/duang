package dockerclienteng1

import (
	"fmt"
	"github.com/cst05001/duang/models"
	"github.com/docker/docker/api/types"
	docker "github.com/fsouza/go-dockerclient"
	"regexp"
)

func (this *DockerClientEng1) Run(unit *models.Unit) error {
	/*
		containerCreateResponse, err := this.CreateContainer(unit)
		if err != nil {
			return err
		}
		err = this.StartContainer(containerCreateResponse.ID, unit)
		return err
	*/
	// create container
	hostConfig := &docker.HostConfig{}
	config := &docker.Config{
		Image: unit.Image,
	}
	createContainerOptions := &docker.CreateContainerOptions{
		Name:       unit.Name,
		Config:     config,
		HostConfig: hostConfig,
	}
	containerCreateResponse := &types.ContainerCreateResponse{}
	for _, client := range this.Client {
		container, err := client.CreateContainer(*createContainerOptions)
		if err != nil {
			fmt.Println(err)
			containerCreateResponse.Warnings = append(containerCreateResponse.Warnings, err.Error())
			return err
		}
		fmt.Println(container)
		containerCreateResponse.ID = container.ID

		// start container
		for _, p := range unit.Parameteres {
			switch p.Type {
			case "v":
				hostConfig.Binds = append(hostConfig.Binds, p.Value)
			case "p":
				rePort := regexp.MustCompile(".+/.+")
				re3 := regexp.MustCompile("(.+):(.+):(.+)")
				if re3.MatchString(p.Value) {
					t := re3.FindStringSubmatch(p.Value)
					portBinding := &docker.PortBinding{
						HostIP:   t[1],
						HostPort: t[2],
					}
					var containerPort string
					if rePort.MatchString(t[3]) {
						containerPort = t[3]
					} else {
						containerPort = fmt.Sprintf("%s/tcp", t[3])
					}
					hostConfig.PortBindings = make(map[docker.Port][]docker.PortBinding)
					hostConfig.PortBindings[docker.Port(containerPort)] = append(hostConfig.PortBindings[docker.Port(containerPort)], *portBinding)
					break
				}
				re2 := regexp.MustCompile("(.+):(.+)")
				if re2.MatchString(p.Value) {
					t := re2.FindStringSubmatch(p.Value)
					portBinding := &docker.PortBinding{
						HostPort: t[1],
					}
					hostConfig.PortBindings = make(map[docker.Port][]docker.PortBinding)
					hostConfig.PortBindings[docker.Port(t[2])] = append(hostConfig.PortBindings[docker.Port(t[2])], *portBinding)
					break
				}
				re1 := regexp.MustCompile("(.+)")
				if re1.MatchString(p.Value) {
					t := re2.FindStringSubmatch(p.Value)
					portBinding := &docker.PortBinding{}
					hostConfig.PortBindings = make(map[docker.Port][]docker.PortBinding)
					hostConfig.PortBindings[docker.Port(t[1])] = append(hostConfig.PortBindings[docker.Port(t[1])], *portBinding)
					break
				}
			}
		}
		fmt.Println(hostConfig.PortBindings)
		err = client.StartContainer(containerCreateResponse.ID, hostConfig)
		if err != nil {
			fmt.Sprintf("StartContainer: Faile with error %s\n", err)
			return err
		}
	}
	return nil
}
