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
	Addr     []*BackendAddr `orm:"reverse(many);on_delete(set_null)"`
	Frontend *Frontend      `orm:"rel(fk)"`
}

type BackendAddr struct {
	Id      int64
	Value   string   `orm:"unique"`
	Backend *Backend `orm:"rel(fk)"`
}
