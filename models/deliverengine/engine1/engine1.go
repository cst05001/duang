package engine1

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/coreos/go-etcd/etcd"
	"path"
	"strings"
)

type Engine1 struct {
	Addr []string
	Root string
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
	etcdClient := etcd.NewClient(this.Addr)
	_, err = EtcdLs(etcdClient, this.Root)
	if err != nil {
		errorCode, _, _ := ParseError(err)
		if errorCode == "100" {
			err = EtcdMkDir(etcdClient, this.Root, 0)
			if err != nil {
				fmt.Println(err)
				return nil
			}
		}
	}
	_, err = EtcdLs(etcdClient, path.Join(this.Root, "backend"))
	if err != nil {
		errorCode, _, _ := ParseError(err)
		if errorCode == "100" {
			err = EtcdMkDir(etcdClient, path.Join(this.Root, "backend"), 0)
			if err != nil {
				fmt.Println(err)
				return nil
			}
		}
	}
	_, err = EtcdLs(etcdClient, path.Join(this.Root, "frontend"))
	if err != nil {
		errorCode, _, _ := ParseError(err)
		if errorCode == "100" {
			err = EtcdMkDir(etcdClient, path.Join(this.Root, "frontend"), 0)
			if err != nil {
				fmt.Println(err)
				return nil
			}
		}
	}
	return nil
}

func (this *Engine1) SetBackend(frontend string, backends []string) error {
	return nil
}
func (this *Engine1) AddBackend(frontend string, backends []string) error {
	return nil
}
func (this *Engine1) DelBackend(frontend string, backends []string) error {
	return nil
}
func (this *Engine1) GetBackend(frontend string) []string {
	return nil
}
func (this *Engine1) GetFrontend() []string {
	return nil
}
func (this *Engine1) PauseFrontend(frontend string) error {
	return nil
}
func (this *Engine1) ResumeFrontend(frontend string) error {
	return nil
}
func (this *Engine1) DelFrontend(frontend string) error {
	return nil
}
