package models

type Unit struct {
	Id          int64
	Name        string
	Image       string
	Number      uint
	Parameteres []string         `orm:"-"`
	Parameter   []*UnitParameter `orm:"reverse(many);on_delete(set_null)"`
}

type UnitParameter struct {
	Id        int64
	Unit      *Unit `orm:"rel(fk)"`
	Parameter string
	Order     int
}
