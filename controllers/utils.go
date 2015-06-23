package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/context"
)

type Status struct {
	Status string
	Msg    string
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
