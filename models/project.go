package models

import "github.com/uadmin/uadmin"

type Project struct {
	uadmin.Model
	Name         string
	Cost         int64
	Phase        []Phase `uadmin:"list_exclude"`
	CreatedBy    uadmin.User
	CreatedByID  uint
	ModifiedBy   uadmin.User
	ModifiedByID uint
}

func (p Project) String() string {
	return p.Name
}

func (p *Project) Save() {
	uadmin.Save(p)
}
