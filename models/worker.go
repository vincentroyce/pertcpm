package models

import "github.com/uadmin/uadmin"

type Worker struct {
	uadmin.Model
	Project      Project
	ProjectID    uint
	Name         string
	RatePerHour  int
	Quantity     int
	Total        int
	CreatedBy    uadmin.User
	CreatedByID  uint
	ModifiedBy   uadmin.User
	ModifiedByID uint
}

func (w Worker) String() string {
	return w.Name
}
