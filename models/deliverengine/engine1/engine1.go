package engine1

import (
	"crypto/md5"
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

func (this *Engine1) Bind(bind, domain string, backend string) error {
	//目录不在则创建目录
	frontendRoot := path.Join(this.Root, "frontend", bind)
	err := MkDirIfNotExist(this.Client, frontendRoot)
	if err != nil {
		return err
	}

	//bind
	_, err = this.Client.Set(path.Join(frontendRoot, "bind"), bind, 0)
	if err != nil {
		return err
	}

	//domain
	domainDir := path.Join(frontendRoot, "domain")
	err = MkDirIfNotExist(this.Client, domainDir)
	if err != nil {
		return err
	}
	backendDir := path.Join(domainDir, "backends")
	err = MkDirIfNotExist(this.Client, backendDir)
	if err != nil {
		return err
	}

	m := fmt.Sprintf("%x", md5.Sum([]byte(backend)))
	_, err = this.Client.Set(path.Join(backendDir, m), backend, 0)
	if err != nil {
		return err
	}

	return nil
}
