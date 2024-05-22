package models

import (
	"strings"

	"github.com/uadmin/uadmin"
)

type SubActivity struct {
	uadmin.Model
	NoPredecessorsList string `uadmin:"read_only"`
	PredecessorsList   string `uadmin:"read_only"`
	Project            Project
	ProjectID          uint
	Phase              Phase
	PhaseID            uint
	Activity           Activity
	ActivityID         uint
	No                 string
	Name               string
	OptimisticTime     int
	MostLikelyTime     int
	PessimisticTime    int
	ExpectedTime       int
	Critical           bool
	PhaseDirect        bool
}

func (s SubActivity) String() string {
	return s.Name
}

func (s *SubActivity) Save() {
	s.ExpectedTime = GetExpectedTime(s.OptimisticTime, s.MostLikelyTime, s.PessimisticTime)

	predlist := strings.Split(s.PredecessorsList, ",")
	prednolist := make([]string, len(predlist))
	for i := range predlist {
		sub := SubActivity{}
		if predlist[i] == "0" {
			prednolist[i] = "Start"
			continue
		}
		uadmin.Get(&sub, "id = ?", predlist[i])
		prednolist[i] = sub.No
	}
	s.NoPredecessorsList = strings.Join(prednolist, ", ")
	uadmin.Save(s)
}
