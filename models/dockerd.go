package models

type Dockerd struct {
	Id   int64
	Addr string  `orm:"unique"`
	Unit []*Unit `orm:"reverse(many)"`
}
