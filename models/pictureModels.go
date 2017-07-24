package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Picture struct {
	Id          int64     `orm:"column(id);pk"                                              json:"id"`
	Created     time.Time `orm:"column(created);type(datetime)"                             json:"created"`
	Updated     time.Time `orm:"column(updated);type(datetime)"                             json:"updated"`
	Url         string    `orm:"column(url);size(255)"                                      json:"url"`
	Feature     string    `orm:"column(feature);size(255)"                                  json:"feature"`
	Library     *Library  `orm:"column(library);size(255);rel(fk)"                          json:"-"`
	Target      *Target   `orm:"column(target);size(255);rel(fk);on_delete(cascade)"        json:"-"`
	LibraryId   int64     `orm:"-"							                   			    json:"libraryId"`
	LibraryName string    `orm:"-"										                    json:"libraryName"`
	TargetId    int64     `orm:"-"										                    json:"targetId"`
	TargetName  string    `orm:"-"										                    json:"targetName"`
	Gender      string    `orm:"-"										                    json:"gender"`
}

type ResPicture struct {
	Id      int64
	Created time.Time
	Updated time.Time
	Url     string
	Feature string
}

type Photo struct {
	Index    int    `json:"index"`
	Name     string `json:"uid"`
	Url      string `json:"url"`
	Feature  string `json:"feature"`
	Verified string `json:"verified"`
}

func init() {
	orm.RegisterModel(new(Picture))
}

func NewPicture() *Picture {
	return &Picture{}
}

func (p *Picture) TableName() string {
	return "pictures"
}

// POST
func (p *Picture) AddPicture() (err error) {
	o := orm.NewOrm()

	p.Created = time.Now()
	p.Updated = time.Now()

	if _, err = o.Insert(p); err != nil {
		logs.Error(err)
		return
	}

	return
}

// UPDATE
func (p *Picture) UpdatePicture() {
}

// DELETE
func (p *Picture) DelPicture() (err error) {
	o := orm.NewOrm()

	if _, err = o.Delete(p); err != nil {
		logs.Error(err)
		return
	}
	return
}

// GET
// Get Picture by any argv
func (p *Picture) LookUp(item map[string]interface{}) (picture Picture, err error) {
	o := orm.NewOrm()

	defer func() {
		if err != nil {
			logs.Error(err)
		}
	}()

	for k, v := range item {
		switch k {
		case "feature":
			if exist := o.QueryTable(new(Picture)).Filter(k, v).Exist(); !exist {
				err = fmt.Errorf("not exist")
				return
			}
			if err = o.QueryTable(new(Picture)).Filter(k, v).One(&picture); err != nil {
				return
			}
		}
	}
	return
}
