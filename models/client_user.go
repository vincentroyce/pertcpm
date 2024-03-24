package models

import "github.com/uadmin/uadmin"

type ClientUser struct {
	uadmin.Model
	Username  string
	Password  string
	FirstName string
	LastName  string
}

func (u *ClientUser) String() string {
	return u.FirstName + " " + u.LastName
}

func (u *ClientUser) Save() {
	uadminUser := uadmin.User{}
	uadminUser.Username = u.Username
	uadminUser.Password = u.Password
	uadminUser.FirstName = u.FirstName
	uadminUser.LastName = u.LastName
	uadminUser.Active = true

	uadminUser.Save()
	uadmin.Save(u)
}
