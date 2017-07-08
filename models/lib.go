package models

import (
	_ "fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"last/errs"
)

type Library struct {
	Id      int64     `orm:"column(id);pk"                           json:"id"`
	Name    string    `orm:"column(name);size(255)"                  json:"name"`
	Created time.Time `orm:"column(created);type(datetime)"          json:"created"`
	Updated time.Time `orm:"column(updated);type(datetime)"          json:"updated"`
	Role    string    `orm:"column(role);size(255)"                  json:"role"`
	Message string    `orm:"column(message);size(255)"               json:"message"`
	Pics    []*Target `orm:"column(Id);reverse(many)"                json:"picture"`
}

type ResLibrary struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Role     string    `json:"role"`
	AimCount int       `json:"aim_count"`
	PicCount int       `json:"pic_count"`
	User     string    `json:"user"`
	Message  string    `json:"message"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}

func init() {
	orm.RegisterModel(new(Library))
}

func NewLibrary() *Library {
	return &Library{}
}

func (l *Library) TableName() string {
	return "librarys"
}

// GET
func GetAllLibrarys() (resLibs []ResLibrary, err error) {
	o := orm.NewOrm()

	defer func() {
		if err != nil {
			logs.Error(err)
			return
		}
	}()

	libs := make([]Library, 0)
	resLibs = make([]ResLibrary, 0)
	if _, err = o.QueryTable(new(Library)).All(&libs); err != nil {
		return
	}

	for _, lib := range libs {
		var res ResLibrary
		res.Id = lib.Id
		res.Name = lib.Name
		res.Role = lib.Role
		res.User = "admin"
		res.Message = lib.Message
		res.Created = lib.Created
		res.Updated = lib.Updated
		if _, err = o.LoadRelated(&lib, "Pics"); err != nil {
			return
		}
		res.AimCount = len(lib.Pics)
		res.PicCount = len(lib.Pics)
		resLibs = append(resLibs, res)
	}

	return
}

// POST
func AddLibrary(name, role, message string) (err error) {
	o := orm.NewOrm()

	var l Library
	l.Name = name
	l.Role = role
	l.Message = message
	l.Created = time.Now()
	l.Updated = time.Now()

	if _, err = o.Insert(&l); err != nil {
		return
	}
	return
}

// UPDATE
func UpdateLibrary(name, role, message string, id int64) (err error) {
	o := orm.NewOrm()

	if err = checkLibraryById(id); err == nil {
		l := Library{Id: id}
		l.Name = name
		l.Role = role
		l.Message = message
		l.Updated = time.Now()
		if _, err = o.Update(&l); err != nil {
			logs.Error(err)
			return
		}
	}
	return
}

// DELETE
func DelLibrary(id int64) (err error) {
	o := orm.NewOrm()

	if err = checkLibraryById(id); err == nil {
		if _, err = o.Raw("delete from librarys where id= ?", id).Exec(); err != nil {
			logs.Error(err)
			return
		}
	}
	return
}

func checkLibraryById(id int64) (err error) {
	o := orm.NewOrm()

	num, err := o.Raw("select * from librarys where id = ?", id).QueryRows(&[]Library{})
	if err != nil {
		return
	} else if num == 0 {
		return errs.LibraryNotFound
	}
	return
}

func GetLibraryById(id int64) (l *Library, err error) {
	o := orm.NewOrm()

	if err = o.Raw("select * from librarys where id = ?", id).QueryRow(&l); err != nil {
		return
	}
	return
}
