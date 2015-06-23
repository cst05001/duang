package core

type Dockerd struct {
	Id   int64
	Addr string  `orm:"unique"`
	Unit []*Unit `orm:"reverse(many)"`
}

type DockerdSchedulerInterface interface {
	GetDockerd(n int64) []*Dockerd
}
