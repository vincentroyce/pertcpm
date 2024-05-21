package models

import (
	"github.com/uadmin/uadmin"
)

type SubActivity struct {
	uadmin.Model
	PredecessorsList string `uadmin:"read_only"`
	Project          Project
	ProjectID        uint
	Phase            Phase
	PhaseID          uint
	Activity         Activity
	ActivityID       uint
	No               string
	Name             string
	OptimisticTime   int
	MostLikelyTime   int
	PessimisticTime  int
	ExpectedTime     int
	Critical         bool
	PhaseDirect      bool
}

func (s SubActivity) String() string {
	return s.Name
}

func (s *SubActivity) Save() {
	s.ExpectedTime = GetExpectedTime(s.OptimisticTime, s.MostLikelyTime, s.PessimisticTime)
	uadmin.Save(s)
}
