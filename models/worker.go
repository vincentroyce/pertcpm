package models

import "github.com/uadmin/uadmin"

type Worker struct {
	uadmin.Model
	Project   Project
	ProjectID uint
	Phase     Phase
	PhaseID   uint
	Name      string
	Rate      int
	Quantity  int
	Total     int
}

func (w Worker) String() string {
	return w.Name
}
