package models

import "github.com/uadmin/uadmin"

type Activity struct {
	uadmin.Model
	Project         Project
	ProjectID       uint
	Phase           Phase
	PhaseID         uint
	Name            string
	OptimisticTime  int
	MostLikelyTime  int
	PessimisticTime int
	ExpectedTime    int
}

func (a Activity) String() string {
	return a.Name
}

func (a *Activity) Save() {
	a.ExpectedTime = GetExpectedTime(a.OptimisticTime, a.MostLikelyTime, a.PessimisticTime)
	uadmin.Save(a)
}
