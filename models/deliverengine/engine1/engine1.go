package engine1

//review at 20150703
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
	_, err = Ls(this.Client, this.Root)
	if err != nil {
		errorCode, _, _ := ParseError(err)
		if errorCode == "100" {
			err = EtcdMkDir(this.Client, this.Root, 0)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	_, err = Ls(this.Client, path.Join(this.Root, "frontend"))
	if err != nil {
		errorCode, _, _ := ParseError(err)
		if errorCode == "100" {
			err = EtcdMkDir(this.Client, path.Join(this.Root, "frontend"), 0)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}
	return nil
}

func (this *Engine1) Bind(bind, domain string, backend string) error {
	var err error

	frontendRoot := path.Join(this.Root, "frontend", bind)

	_, err = this.Client.Set(path.Join(frontendRoot, "bind"), bind, 0)
	if err != nil {
		fmt.Println("debug 2")
		return err
	}

	backendDir := path.Join(frontendRoot, "domain", domain, "backends")

	m := fmt.Sprintf("%x", md5.Sum([]byte(backend)))
	_, err = this.Client.Set(path.Join(backendDir, m), backend, 0)
	if err != nil {
		fmt.Println("debug 6")
		return err
	}

	return nil
}
