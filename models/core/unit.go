package core

type Unit struct {
	Id          int64
	Name        string `orm:"unique"`
	Image       string
	Number      int64
	Domain      string
	Status      uint8            `orm:default(0)`
	Parameteres []*UnitParameter `orm:"reverse(many)"	json:"-"`
	Dockerd     []*Dockerd       `orm:"rel(m2m)"	json:"-"`
}

type UnitParameter struct {
	Id    int64
	Unit  *Unit `orm:"rel(fk)"	json:"-"`
	Value string
	Type  string
}
