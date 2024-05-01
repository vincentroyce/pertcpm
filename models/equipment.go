package models

import "github.com/uadmin/uadmin"

type Equipment struct {
	uadmin.Model
	Project     Project
	ProjectID   uint
	Name        string
	CostPerHour int
	Quantity    int
	Total       int
}

func (e Equipment) String() string {
	return e.Name
}
