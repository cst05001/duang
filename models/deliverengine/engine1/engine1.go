package engine1

import (
	"strings"
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/coreos/go-etcd/etcd"
)
type Engine1 struct {
	Addr	[]string
	Root 	string
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
    response , err := etcdClient.Get(this.Root, true, false)
    if err != nil {
    	errorCode, _, _ := ParseError(err)
    	if errorCode == "100" {
    		_, err = etcdClient.SetDir(this.Root, 0)
    		if err != nil {
    			fmt.Println(err)
    			return nil
    		}
    	}
    } else {
    	fmt.Println(response.Node)
    }
    //测试代码
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