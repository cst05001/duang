package engine1

import (
	"github.com/coreos/go-etcd/etcd"
	"regexp"
)

func ParseError(err error) (code, text, index string) {
	re := regexp.MustCompile("^\\s*(.+)\\s*:\\s*(.+)\\s*\\[(.+)\\]\\s*$")

	if re.MatchString(err.Error()) {
		result := re.FindStringSubmatch(err.Error())
		code = result[1]
		text = result[2]
		index = result[3]
		return
	}
	return "", "", ""
}

func MkDirIfNotExist(client *etcd.Client, path string) error {
	_, err := EtcdLs(client, path)
	if err != nil {
		errorCode, _, _ := ParseError(err)
		if errorCode == "100" {
			err = EtcdMkDir(client, path, 0)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return nil
}

func EtcdLs(client *etcd.Client, path string) ([]string, error) {
	result := make([]string, 0)
	response, err := client.Get(path, true, false)
	if err != nil {
		return nil, err
	}

	//文件，列出自身路径
	if response.Node.Nodes == nil {
		return append(result, response.Node.Value), nil
	}

	//目录，列出目录内容
	for _, n := range response.Node.Nodes {
		result = append(result, n.Value)

	}
	return result, nil
}

func EtcdMkDir(client *etcd.Client, path string, ttl uint64) error {
	_, err := client.SetDir(path, ttl)
	if err != nil {
		return err
	}
	return nil
}
