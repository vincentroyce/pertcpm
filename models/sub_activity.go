package models

import "github.com/uadmin/uadmin"

type SubActivity struct {
	uadmin.Model
	Key             int
	Project         Project
	ProjectID       uint
	Phase           Phase
	PhaseID         uint
	Activity        Activity
	ActivityID      uint
	Name            string
	OptimisticTime  int
	MostLikelyTime  int
	PessimisticTime int
	ExpectedTime    int
	Critical        bool
	PhaseDirect     bool
}

func (s SubActivity) String() string {
	return s.Name
}

func (s *SubActivity) Save() {
	s.ExpectedTime = GetExpectedTime(s.OptimisticTime, s.MostLikelyTime, s.PessimisticTime)
}
