package engine1

//reviewed at 20150703
import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

type SshClient struct {
	Client  *ssh.Client
	Session *ssh.Session
}

func (this *SshClient) Start(cmd string) error {
	err := this.Session.Start(cmd)
	if err != nil {
		beego.Error("Sshclient start cmd: ", cmd, " failed at ", this.Client.RemoteAddr().String())
		return err
	}
	beego.Debug("Sshclient start cmd: ", cmd, " successed at ", this.Client.RemoteAddr().String())
	return nil
}

func (this *SshClient) Run(cmd string) error {
	err := this.Session.Run(cmd)
	if err != nil {
		beego.Error("Sshclient run cmd: ", cmd, " failed at ", this.Client.RemoteAddr().String())
		return err
	}
	beego.Debug("Sshclient run cmd: ", cmd, " successed at ", this.Client.RemoteAddr().String())
	return nil
}

func (this *SshClient) Output(cmd string) ([]byte, error) {
	response, err := this.Session.Output(cmd)
	if err != nil {
		beego.Error("Sshclient output cmd: ", cmd, " failed at ", this.Client.RemoteAddr().String())
		return nil, err
	}
	beego.Debug("Sshclient output cmd: ", cmd, " successed at ", this.Client.RemoteAddr().String())
	return response, nil
}

func (this *SshClient) CombinedOutput(cmd string) ([]byte, error) {
	response, err := this.Session.CombinedOutput(cmd)
	if err != nil {
		beego.Error("Sshclient combined output cmd: ", cmd, " failed at ", this.Client.RemoteAddr().String())
		return nil, err
	}
	beego.Debug("Sshclient combined output cmd: ", cmd, " successed at ", this.Client.RemoteAddr().String())
	return response, nil
}

func (this *SshClient) Close() error {
	err := this.Session.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = this.Client.Close()
	if err != nil {
		return err
	}
	return nil
}

func NewSSLClient(addr, user, keypath string) (*SshClient, error) {
	duangcfg, err := config.NewConfig("ini", "conf/duang.conf")
	if err != nil {
		return nil, err
	}

	_ = duangcfg.String("pipework_path")

	key, err := ioutil.ReadFile(keypath)
	if err != nil {
		return nil, err
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	}

	sshClient := &SshClient{}
	sshClient.Client, err = ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		beego.Error("Sshclient dail to ", addr, " failed: ", err)
		return nil, err
	}
	sshClient.Session, err = sshClient.Client.NewSession()
	if err != nil {
		beego.Error("Sshclient new session failed at ", addr, ": ", err)
		return nil, err
	}
	return sshClient, nil
}
