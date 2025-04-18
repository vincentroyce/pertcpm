package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Project struct {
	uadmin.Model
	Name              string
	Description       string
	Cost              int64   `uadmin:"money"`
	Phase             []Phase `uadmin:"list_exclude"`
	DateStart         time.Time
	DateEnd           *time.Time
	CompletedAt       *time.Time
	ScheduleCompleted bool
	Completed         bool
	Reply             string `uadmin:"link"`
}

func (p Project) String() string {
	return p.Name
}

func (p *Project) Save() {
	p.Reply = "https://www.google.com"
	uadmin.Save(p)
}
