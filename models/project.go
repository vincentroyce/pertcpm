package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Project struct {
	uadmin.Model
	Name              string
	Description       string
	Cost              int64
	Phase             []Phase `uadmin:"list_exclude"`
	DateStart         time.Time
	DateEnd           time.Time
	ExpiresAt         *time.Time
	ScheduleCompleted bool
	CreatedBy         uadmin.User
	CreatedByID       uint
	ModifiedBy        uadmin.User
	ModifiedByID      uint
	Completed         bool
}

func (p Project) String() string {
	return p.Name
}

func (p *Project) Save() {
	uadmin.Save(p)
}
