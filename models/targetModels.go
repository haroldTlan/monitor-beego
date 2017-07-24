package models

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/lifei6671/mindoc/conf"
	"github.com/lifei6671/mindoc/errs"
	"github.com/lifei6671/mindoc/utils"
)

type Target struct {
	Id          int64      `orm:"column(id);pk"                     json:"id"`
	Name        string     `orm:"column(name);size(255)"            json:"name"`
	Created     time.Time  `orm:"column(created);type(datetime)"    json:"created"`
	Updated     time.Time  `orm:"column(updated);type(datetime)"    json:"updated"`
	Identity    string     `orm:"column(identity);size(255)"        json:"identity"`
	Gender      string     `orm:"column(sex);size(255)"             json:"gender"`
	Level       int64      `orm:"column(level)"                     json:"level"`
	Age         int64      `orm:"column(age)"                       json:"age"`
	Nation      string     `orm:"column(nation);size(255)"          json:"nation"`
	Host        string     `orm:"column(host);size(255)"            json:"host"`
	Message     string     `orm:"column(message);size(255)"         json:"message"`
	Library     *Library   `orm:"column(library);size(255);rel(fk)" `
	Pictures    []*Picture `orm:"column(Id);reverse(many)"          json:"pictures"`
	LibraryName string     `orm:"-"                                 json:"libraryName"`
	LibraryId   int64      `orm:"-"                                 json:"libraryId"`
	Url         string     `orm:"-"                                 json:"url"`
	/*
		LibraryId int64     `orm:"-"                                 json:"libraryId"`
		LibraryName string     `orm:"-"                                 json:"libraryName"`
	*/
}

type ResTarget struct {
	Id       int64     `json:"id"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
	Name     string    `json:"name"`
	Identity string    `json:"identity"`
	Url      string    `json:"url"`
	Gender   string    `json:"gender"`
	Level    int64     `json:"level"`
	Age      int64     `json:"age"`
	Nation   string    `json:"nation"`
	Host     string    `json:"host"`
	Message  string    `json:"message"`
	Library  string    `json:"library"`
}

type Feature struct {
	FileId string `json:"fid"`
}

type DataPic struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Feature string `json:"feature"`
}

var (
	InstallFilePath  = beego.AppConfig.String("InstallFilePath")
	WorkingDirectory = ""
)

func init() {
	orm.RegisterModel(new(Target))

	if WorkingDirectory == "" {
		if p, err := filepath.Abs(os.Args[0]); err == nil {
			WorkingDirectory = filepath.Dir(p)
		}
	}
}

func NewTarget() *Target {
	return &Target{}
}

func (t *Target) TableName() string {
	return "targets"
}

// GET
func (t *Target) GetTargetByLib(libId int64) (ts []Target, err error) {
	o := orm.NewOrm()

	ts = make([]Target, 0)
	if _, err = o.QueryTable(new(Target)).Filter("library", libId).All(&ts); err != nil {
		return
	}

	for i, t := range ts {
		if t.Library == nil {
			t.LibraryName = ""
			err = errs.NotInAnyLibrary
			return
		} else {
			if _, err = o.LoadRelated(&t, "Library"); err != nil {
				return
			}
			t.LibraryName = t.Library.Name
		}
		if _, err = o.LoadRelated(&t, "Pictures"); err != nil {
			return
		}

		if len(t.Pictures) > 0 {
			if _, err = os.Stat(conf.InstallFilePath + "/mindoc" + t.Pictures[0].Url); err == nil {
				t.Url = t.Pictures[0].Url
			} else {
				t.Url = "/uploads/none.svg"
			}
		} else {
			t.Url = "/uploads/none.svg"
		}
		ts[i] = t
	}
	return
}

// POST
func (t *Target) AddTarget(photos []Photo) (err error) {
	o := orm.NewOrm()

	defer func() {
		if err != nil {
			logs.Error(err)
		}
	}()

	t.Created = time.Now()
	t.Updated = time.Now()
	if t.Library, err = NewLibrary().CheckLibraryById(t.LibraryId); err != nil {
		return
	}

	if _, err = o.Insert(t); err != nil {
		return
	}

	for _, photo := range photos {
		p := NewPicture()
		p.Url = photo.Url
		p.Feature = photo.Feature
		target, _ := t.LookUp(map[string]interface{}{"name": t.Name})
		p.Target = &target
		p.Library = t.Library
		p.AddPicture()
	}
	return
}

// GET
// Get Target by any argv
func (t *Target) LookUp(item map[string]interface{}) (target Target, err error) {
	o := orm.NewOrm()

	defer func() {
		if err != nil {
			logs.Error(err)
		}
	}()

	for k, v := range item {
		switch k {
		case "name":
			if exist := o.QueryTable(new(Target)).Filter(k, v).Exist(); !exist {
				err = fmt.Errorf("not exist")
				return
			}
			if err = o.QueryTable(new(Target)).Filter(k, v).One(&target); err != nil {
				return
			}
		case "id":
			if exist := o.QueryTable(new(Target)).Filter(k, v).Exist(); !exist {
				err = fmt.Errorf("not exist")
				return
			}
			if err = o.QueryTable(new(Target)).Filter(k, v).One(&target); err != nil {
				return
			}

		}
	}

	if target.Library == nil {
		target.LibraryName = ""
		err = errs.NotInAnyLibrary
		return
	} else {
		if _, err = o.LoadRelated(&target, "Library"); err != nil {
			return
		}
		target.LibraryName = target.Library.Name
		target.LibraryId = target.Library.Id
	}
	if _, err = o.LoadRelated(&target, "Pictures"); err != nil {
		return
	}
	return
}

// POST
// UPDATE
func (t *Target) UpdateTarget(photos []Photo) (err error) {
	o := orm.NewOrm()

	defer func() {
		if err != nil {
			logs.Error(err)
		}
	}()

	if _, err = o.LoadRelated(t, "Library"); err != nil {
		return
	}
	if _, err = o.LoadRelated(t, "Pictures"); err != nil {
		return
	}

	pictureDB := make(map[string]bool, 0)
	for _, p := range t.Pictures {
		pictureDB[p.Feature] = false
	}

	for _, p := range photos {
		if _, ok := pictureDB[p.Feature]; ok {
			pictureDB[p.Feature] = true
		} else {
			fmt.Println("create New:", p.Feature)
			newP, err := NewPicture().LookUp(map[string]interface{}{"feature": p.Feature})
			fmt.Println(err)
			newP.Url = p.Url
			newP.Feature = p.Feature
			target, _ := t.LookUp(map[string]interface{}{"name": t.Name})
			newP.Target = &target
			newP.Library = t.Library
			newP.AddPicture()
		}
	}

	for k, v := range pictureDB {
		if !v {
			oldP, err := NewPicture().LookUp(map[string]interface{}{"feature": k})
			fmt.Println("delete:", oldP.Url)
			if err != nil {
				return err
			}
			if _, err = o.QueryTable(new(Similar)).Filter("picture_id", oldP.Id).Delete(); err != nil {
				return err
			}
			oldP.DelPicture()
		}
	}

	t.Updated = time.Now()

	if _, err = o.Update(t); err != nil {
		return
	}

	return
}

// DELETE
func (t *Target) DelTarget() (err error) {
	o := orm.NewOrm()

	fmt.Printf("%+v", t)
	if _, err = o.LoadRelated(t, "Pictures"); err != nil {
		return
	}

	for _, p := range t.Pictures {
		if _, err = o.QueryTable(new(Similar)).Filter("picture_id", p.Id).Delete(); err != nil {
			return
		}
		p.DelPicture()
	}

	if _, err = o.Delete(t); err != nil {
		logs.Error(err)
		return
	}

	return
}

func uploadFiles(file multipart.File, path string) (err error) {
	defer func() {
		if err != nil {
			logs.Error(err)
		}
	}()

	img, err := os.Create(path)
	defer img.Close()
	if err != nil {
		return
	}

	if _, err = io.Copy(img, file); err != nil {
		return
	}
	return
}

func GetFeature() (res Feature, err error) {
	path := "/tmp/feat.jpg"

	var cmd []string
	cmd = append(cmd, "getfeature.py", path)
	filepath := InstallFilePath
	name := "python"

	o, err := utils.Execute(filepath, name, cmd, true)
	if err = json.Unmarshal([]byte(o), &res); err != nil {
		return res, fmt.Errorf("not vaild picture")
	}

	return
}

func GetPictureByFeature(fid string) (pic Picture, err error) {
	o := orm.NewOrm()
	//var pic Picture
	if _, err = o.QueryTable(new(Picture)).Filter("feature", fid).All(&pic); err != nil {
	}

	if _, err = o.LoadRelated(&pic, "Target"); err != nil {
	}

	if _, err = o.LoadRelated(&pic, "Library"); err != nil {
	}

	pic.LibraryId = pic.Library.Id
	pic.LibraryName = pic.Library.Name
	pic.Gender = pic.Target.Gender
	pic.TargetId = pic.Target.Id
	pic.TargetName = pic.Target.Name

	return
}

func CopyFileOnFacial(filetype string, fileHeader *multipart.FileHeader, feature Feature) (d DataPic, err error) {
	var fName string
	if filetype == "success" {
		fName = strings.Replace(feature.FileId, ",", "-", -1) + ".jpg"
		d.Feature = feature.FileId
	} else {
		fName = strings.Replace(time.Now().String(), " ", "", -1) + ".jpg"
		d.Feature = ""
	}

	url := filepath.Join(WorkingDirectory, "uploads", "facial", fName)

	//fi, err := os.OpenFile(url, os.O_CREATE|os.O_WRONLY, 0755)
	file, err := fileHeader.Open()
	defer file.Close()
	if err != nil {
	}
	fi, err := os.Create(url)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	if _, err := io.Copy(fi, file); err != nil {
	}

	d.Name = fName
	d.Url = "/uploads/facial/" + fName

	return
}
