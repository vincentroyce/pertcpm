package models

import "github.com/uadmin/uadmin"

type Project struct {
	uadmin.Model
	Name         string
	Cost         int64
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
