package engine1

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/coreos/go-etcd/etcd"
	"path"
	"strings"
)

type Engine1 struct {
	Addr   []string
	Root   string
	Client *etcd.Client
}

func NewDeliver() *Engine1 {
	e := &Engine1{}
	e.Init()
	return e
}

func (this *Engine1) Init() error {
	duangcfg, err := config.NewConfig("ini", "conf/duang.conf")
	if err != nil {
		fmt.Println(err)
		return err
	}
	etcd_addr := duangcfg.String("etcd_addr")
	this.Root = duangcfg.String("etcd_root")

	this.Addr = strings.Split(etcd_addr, ",")

	//如果 root 不存在，则创建。
	this.Client = etcd.NewClient(this.Addr)
	_, err = EtcdLs(this.Client, this.Root)
	if err != nil {
		errorCode, _, _ := ParseError(err)
		if errorCode == "100" {
			err = EtcdMkDir(this.Client, this.Root, 0)
			if err != nil {
				fmt.Println(err)
				return nil
			}
		}
	}
	_, err = EtcdLs(this.Client, path.Join(this.Root, "backend"))
	if err != nil {
		errorCode, _, _ := ParseError(err)
		if errorCode == "100" {
			err = EtcdMkDir(this.Client, path.Join(this.Root, "backend"), 0)
			if err != nil {
				fmt.Println(err)
				return nil
			}
		}
	}
	_, err = EtcdLs(this.Client, path.Join(this.Root, "frontend"))
	if err != nil {
		errorCode, _, _ := ParseError(err)
		if errorCode == "100" {
			err = EtcdMkDir(this.Client, path.Join(this.Root, "frontend"), 0)
			if err != nil {
				fmt.Println(err)
				return nil
			}
		}
	}
	return nil
}

//Backend
func (this *Engine1) AddBackend(frontend string, backends []string) error {
	return nil
}
func (this *Engine1) DelBackend(frontend string, backends []string) error {
	return nil
}

//Frontend
func (this *Engine1) AddFrontend(frontend string) error {
	frontendRoot := path.Join(this.Root, "frontend", frontend)
	EtcdMkDir(this.Client, frontendRoot, 0)
	_, err := this.Client.Set(path.Join(frontendRoot, "name"), frontend, 0)
	if err != nil {
		return err
	}
	return nil
}
func (this *Engine1) DelFrontend(frontend string) error {
	return nil
}
