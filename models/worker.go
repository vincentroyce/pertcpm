package models

import "github.com/uadmin/uadmin"

type Worker struct {
	uadmin.Model
	Project   Project
	ProjectID uint
	Phase     Phase
	PhaseID   uint
	Name      string
	Rate      int `uadmin:"money"`
	Quantity  int
	Total     int `uadmin:"money"`
}

func (w Worker) String() string {
	return w.Name
}

func (w *Worker) Save() {
	w.Total = w.Rate * w.Quantity
	uadmin.Save(w)
}
