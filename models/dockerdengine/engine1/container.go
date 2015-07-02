package engine1

import (
	"fmt"
	"regexp"

	"github.com/cst05001/duang/models/core"
	"github.com/cst05001/duang/models/dockerdengine"
	"github.com/docker/docker/api/types"
	docker "github.com/fsouza/go-dockerclient"
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

	/*
		解析Unit参数
		Config 是 create 所需的参数
		HostConfig 是 run 所需的参数
		这是 go-dockerclient 所定义的
		参考：https://github.com/Tonnu/go-dockerclient/blob/master/container.go

		当前用单线程来和所有的dockerd交互，到时候要改成携程。
	*/
	for _, p := range unit.Parameteres {
		switch p.Type {
		case "v": //-v Volume
			hostConfig.Binds = append(hostConfig.Binds, p.Value)
		case "p": //-p EXPOSE
			rePort := regexp.MustCompile(".+/.+")
			// 127.0.0.1:80:8080
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
			//80:8080
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

	//强行从registry pull最新版本的image
	pullImageOptions := docker.PullImageOptions{}
	// registry.ws.com/cst05001/nginx:latest
	reImage3 := regexp.MustCompile("^(.*\\.\\w+)(/.*):(.*)")
	// /cst05001/nginx:latest
	reImage2 := regexp.MustCompile("^(/.*):(.*)")
	if reImage3.MatchString(unit.Image) {
		result := reImage3.FindStringSubmatch(unit.Image)
		pullImageOptions.Registry = result[1]
		pullImageOptions.Repository = result[2]
		pullImageOptions.Tag = result[3]
	} else if reImage2.MatchString(unit.Image) {
		result := reImage2.FindStringSubmatch(unit.Image)
		pullImageOptions.Registry = "http://docker.io"
		pullImageOptions.Repository = result[1]
		pullImageOptions.Tag = result[2]
	} else {
		pullImageOptions.Registry = "http://docker.io"
		pullImageOptions.Repository = unit.Image
		pullImageOptions.Tag = "latest"
	}

	for ptrDockerd, ptrClient := range this.ClientMap {
		dockerd := &(*ptrDockerd)
		client := &(*ptrClient)
		go func() {
			//第二个参数支持registry身份认证，还没处理。
			err := client.PullImage(pullImageOptions, docker.AuthConfiguration{})
			if err != nil {
				fmt.Println(err)
				return
			}

			container, err := client.CreateContainer(*createContainerOptions)
			if err != nil {
				fmt.Println(err)
				containerCreateResponse.Warnings = append(containerCreateResponse.Warnings, err.Error())
				callbackFunc(dockerd, dockerdengine.STATUS_ON_CREATE_FAILED, unit)
				return
			}
			fmt.Println(container)
			containerCreateResponse.ID = container.ID
			callbackFunc(dockerd, dockerdengine.STATUS_ON_CREATE_SUCCESSED, unit)

			// start container

			fmt.Println(hostConfig.PortBindings)
			err = client.StartContainer(containerCreateResponse.ID, hostConfig)
			if err != nil {
				fmt.Sprintf("StartContainer: Faile with error %s\n", err)
				callbackFunc(dockerd, dockerdengine.STATUS_ON_RUN_FAILED, unit)
				return
			}
			callbackFunc(dockerd, dockerdengine.STATUS_ON_RUN_SUCCESSED, unit)
		}()
	}
	return nil
}
