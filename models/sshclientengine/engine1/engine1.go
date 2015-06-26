package engine1

import (
	"fmt"
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
		return err
	}
	return nil
}

func (this *SshClient) Run(cmd string) error {
	err := this.Session.Run(cmd)
	if err != nil {
		return err
	}
	return nil
}

func (this *SshClient) Output(cmd string) ([]byte, error) {
	response, err := this.Session.Output(cmd)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (this *SshClient) CombinedOutput(cmd string) ([]byte, error) {
	response, err := this.Session.CombinedOutput(cmd)
	if err != nil {
		return nil, err
	}
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
		fmt.Println(err)
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
		return nil, err
	}
	sshClient.Session, err = sshClient.Client.NewSession()
	if err != nil {
		return nil, err
	}
	return sshClient, nil
}
