package models

import "github.com/uadmin/uadmin"

type Equipment struct {
	uadmin.Model
	Project      Project
	ProjectID    uint
	Name         string
	CostPerHour  int
	Quantity     int
	Total        int
	CreatedBy    uadmin.User
	CreatedByID  uint
	ModifiedBy   uadmin.User
	ModifiedByID uint
}

func (e Equipment) String() string {
	return e.Name
}
