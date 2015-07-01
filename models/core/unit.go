package core

type Unit struct {
	Id          int64
	Name        string `orm:"unique"`
	Image       string
	Number      int64
	Domain      string
	Parameteres []*UnitParameter `orm:"reverse(many)"`
	Dockerd     []*Dockerd       `orm:"rel(m2m)"`
}

type UnitParameter struct {
	Id    int64
	Unit  *Unit `orm:"rel(fk)"`
	Value string
	Type  string
}
