package engine1

import (
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	"github.com/coreos/go-etcd/etcd"
	"github.com/cst05001/duang/models/core"
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
func (this *Engine1) AddBackend(name string, backend *core.Backend) (*core.Backend, error) {
	o := orm.NewOrm()
	o.Using("default")

	var err error
	backend.Id, err = o.Insert(backend)
	if err != nil {
		return nil, err
	}
	err = o.Read(backend)
	if err != nil {
		return nil, err
	}

	//目录不在则创建目录
	backendRoot := path.Join(this.Root, "backend", backend.Name)
	err = MkDirIfNotExist(this.Client, backendRoot)
	if err != nil {
		return nil, err
	}

	serverDir := path.Join(backendRoot, "server")
	err = MkDirIfNotExist(this.Client, serverDir)
	if err != nil {
		return nil, err
	}

	_, err = this.Client.Set(path.Join(serverDir, backend.Name), backend.Addr, 0)
	if err != nil {
		return nil, err
	}

	return backend, nil
}
func (this *Engine1) DelBackend(name string, backend *core.Backend) error {
	return nil
}

//Frontend
func (this *Engine1) AddFrontend(frontend *core.Frontend) (*core.Frontend, error) {
	o := orm.NewOrm()
	o.Using("default")

	var err error
	frontend.Id, err = o.Insert(frontend)
	if err != nil {
		return nil, err
	}
	err = o.Read(frontend)
	if err != nil {
		return nil, err
	}

	//目录不在则创建目录
	frontendRoot := path.Join(this.Root, "frontend", frontend.Name)
	err = MkDirIfNotExist(this.Client, frontendRoot)
	if err != nil {
		return nil, err
	}

	//bind
	EtcdMkDir(this.Client, frontendRoot, 0)
	_, err = this.Client.Set(path.Join(frontendRoot, "bind"), frontend.Bind, 0)
	if err != nil {
		return nil, err
	}
	return frontend, nil
}
func (this *Engine1) DelFrontend(frontend *core.Frontend) error {
	return nil
}
