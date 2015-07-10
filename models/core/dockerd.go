package core

import (
	"regexp"
)

//review at 20150702

type Dockerd struct {
	Id   int64
	Addr string `orm:"unique"`
	//Unit      []*Unit    `orm:"reverse(many)"`
	Container *Container `orm:"reverse(one)"	json:"-"`
}

func (this *Dockerd) GetIP() string {
	re := regexp.MustCompile("^(http|https)://(\\d+\\.\\d+\\.\\d+\\.\\d+):(\\d+)[/]?$")
	result := re.FindStringSubmatch(this.Addr)
	return result[2]
}
