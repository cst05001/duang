package core

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
