package models

import "github.com/uadmin/uadmin"

type Phase struct {
	uadmin.Model
	Project         Project
	ProjectID       uint
	Name            string
	OptimisticTime  int
	MostLikelyTime  int
	PessimisticTime int
	ExpectedTime    int
}

func (p Phase) String() string {
	return p.Name
}

func (p *Phase) Save() {
	p.ExpectedTime = GetExpectedTime(p.OptimisticTime, p.MostLikelyTime, p.PessimisticTime)
	uadmin.Save(p)
}
