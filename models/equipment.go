package models

import "github.com/uadmin/uadmin"

type Equipment struct {
	uadmin.Model
	Project   Project
	ProjectID uint
	Phase     Phase
	PhaseID   uint
	Name      string
	Cost      int
	Quantity  int
	Total     int
}

func (e Equipment) String() string {
	return e.Name
}

func (e *Equipment) Save() {
	e.Total = e.Cost * e.Quantity
	uadmin.Save(e)
}
