package sshclientengine

type SshClient interface {
	Start(cmd string) error
	Run(cmd string) error
	Output(cmd string) ([]byte, error)
	CombinedOutput(cmd string) ([]byte, error)
	Close() error
}
