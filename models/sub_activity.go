package models

import (
	"strconv"
	"strings"

	"github.com/uadmin/uadmin"
)

type SubActivity struct {
	uadmin.Model
	Predecessors     []SubActivity `gorm:"-" uadmin:"list_exclude"`
	PredecessorsList string        `uadmin:"read_only"`
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

	predecessorsList := []string{}

	// Append every element to the categoryList array
	for i := range s.Predecessors {
		predecessorsList = append(predecessorsList, strconv.FormatUint(uint64(s.Predecessors[i].ID), 10))
	}

	// Concatenate the categoryList to a single string separated by comma
	joinList := strings.Join(predecessorsList, ", ")

	// Store the joined string to the CategoryList field
	s.PredecessorsList = joinList

	uadmin.Save(s)
}
