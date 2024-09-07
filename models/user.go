package models

import "github.com/beego/beego/v2/client/orm"

type Users struct {
	Id    int    `orm:"auto"`
	Name  string `orm:"size(100)"`
	Email string `orm:"type(text)"`
}

func init() {
	// register model with beego orm
	orm.RegisterModel(new(Users))
}
