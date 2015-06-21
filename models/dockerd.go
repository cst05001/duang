package models

type Dockerd struct {
	Id   int64
	Addr string
	Unit []*Unit `orm:"reverse(many)"`
}
