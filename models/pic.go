package models

import (
	_ "errors"
	"fmt"
	_ "strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type Picture struct {
	Id       int64     `orm:"column(id);pk"`
	Name     string    `orm:"column(name);size(255)"`
	Created  time.Time `orm:"column(created);type(datetime)"`
	Updated  time.Time `orm:"column(updated);type(datetime)"`
	Identity string    `orm:"column(identity);size(255)"`
	Url      string    `orm:"column(url);size(255)"`
	Sex      string    `orm:"column(sex);size(255)"`
	Level    int64     `orm:"column(level)"`
	Age      int64     `orm:"column(age)"`
	Nation   string    `orm:"column(nation);size(255)"`
	Host     string    `orm:"column(host);size(255)"`
	Message  string    `orm:"column(message);size(255)"`
	Library  *Library  `orm:"column(library);size(255);rel(fk)"`
}

type ResPicture struct {
}

func init() {
	orm.RegisterModel(new(Picture))
}

func (p *Picture) TableName() string {
	return "pictures"
}

// GET
func GetAllPictures() (resPics []ResPicture, err error) {
	o := orm.NewOrm()

	var pics []Picture
	if _, err = o.QueryTable(new(Picture)).All(&pics); err != nil {
		fmt.Println(err)
		return
	}

	return
}
