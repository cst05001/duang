package core

//review at 20150702

/*
这段 struct 已经弃用，用 Deliver 替代。
type Frontend struct {
	Id      int64
	Name    string
	Bind    string
	Backend []*Backend `orm:"reverse(many);on_delete(set_null)"`
}

type Backend struct {
	Id       int64
	Name     string
	Addr     string
	Frontend *Frontend `orm:"rel(fk)"`
}
*/
type Deliver struct {
	Id           int64
	Unit         *Unit `orm:"rel(fk)"`
	FrontendPort uint
	BackendPort  uint
}
