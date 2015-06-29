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
func (this *Engine1) AddBackend(name string, backends []string) error {
	return nil
}
func (this *Engine1) DelBackend(name string, backends []string) error {
	return nil
}

//Frontend
func (this *Engine1) AddFrontend(name, bind string) error {
	//目录不在则创建目录
	frontendRoot := path.Join(this.Root, "frontend", name)
	_, err := EtcdLs(this.Client, frontendRoot)
	if err != nil {
		errorCode, _, _ := ParseError(err)
		if errorCode == "100" {
			err = EtcdMkDir(this.Client, frontendRoot, 0)
			if err != nil {
				fmt.Println(err)
				return nil
			}
		}
	}

	//bind
	EtcdMkDir(this.Client, frontendRoot, 0)
	_, err = this.Client.Set(path.Join(frontendRoot, "bind"), bind, 0)
	if err != nil {
		return err
	}
	return nil
}
func (this *Engine1) DelFrontend(name, bind string) error {
	return nil
}
