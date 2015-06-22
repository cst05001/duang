package core

type Dockerd struct {
	Id   int64
	Addr string  `orm:"unique"`
	Unit []*Unit `orm:"reverse(many)"`
}

type DockerdScheduler interface {
	GetDockerd(n int64) []*Dockerd
}
