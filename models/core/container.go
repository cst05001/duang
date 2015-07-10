package core

type Container struct {
	Id          int64
	Dockerd     *Dockerd `orm:"rel(one)"	json:"-"`
	Unit        *Unit    `orm:"rel(fk)"	json:"-"`
	Ip          *Ip      `orm:"rel(one)"	json:"-"`
	ContainerId string
}
