package engine1

import (
	"fmt"
	"github.com/cst05001/duang/models/core"
	"github.com/cst05001/duang/models/dockerdengine"
	"github.com/docker/docker/api/types"
	docker "github.com/fsouza/go-dockerclient"
	"regexp"
)

func (this *DockerClientEng1) Run(unit *core.Unit, callbackFunc func(*core.Dockerd, int, ...interface{})) error {

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
	for dockerd, client := range this.ClientMap {
		container, err := client.CreateContainer(*createContainerOptions)
		if err != nil {
			fmt.Println(err)
			containerCreateResponse.Warnings = append(containerCreateResponse.Warnings, err.Error())
			callbackFunc(dockerd, dockerdengine.STATUS_ON_CREATE_FAILED, unit)
			continue
		}
		fmt.Println(container)
		containerCreateResponse.ID = container.ID
		callbackFunc(dockerd, dockerdengine.STATUS_ON_CREATE_SUCCESSED, unit)

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
			callbackFunc(dockerd, dockerdengine.STATUS_ON_RUN_FAILED, unit)
			continue
		}
		callbackFunc(dockerd, dockerdengine.STATUS_ON_RUN_SUCCESSED, unit)
	}
	return nil
}
