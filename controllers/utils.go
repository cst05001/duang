package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/cst05001/duang/models/core"
)

type StatusError struct {
	Error string
}

func WriteJson(ctx *context.Context, i interface{}) error {
	data, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
		return err
	}
	ctx.WriteString(string(data))
	return nil
}

type ContainersStatus struct {
	Dockerd *core.Dockerd
	Status  uint8
}
