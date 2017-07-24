package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"os"
	"strconv"

	"github.com/astaxie/beego/logs"
	"github.com/lifei6671/mindoc/controllers/web"
	"github.com/lifei6671/mindoc/models"
)

// Operations about Targets
type TargetController struct {
	ManagerController
}

// URLMapping ...
func (t *TargetController) URLMapping() {
	t.Mapping("GetAll", t.Targets)
	t.Mapping("Post", t.AddTarget)
	t.Mapping("Post", t.UpdateTarget)
	t.Mapping("Post", t.DelTarget)
}

// @Title GetAll
// @Description get all Targets
// @Success 200 {object} models.Target
// @router / [get]
func (t *TargetController) Targets() {
	t.TplName = "manager/targets.tpl"

	libId, err := t.GetInt64(":lib")
	lib, _ := models.NewLibrary().CheckLibraryById(libId)

	targets, err := models.NewTarget().GetTargetByLib(libId)
	b, err := json.Marshal(targets)

	if err != nil {
		t.Data["Results"] = template.JS("[]")
	} else {
		t.Data["Results"] = template.JS(string(b))
	}

	t.Data["Targets"] = targets
	t.Data["Lib"] = lib
}

// @Title Get Targets  Json
// @Description get all Targets
// @Success 200 {object} models.Target
// @router / [get]
func (t *TargetController) TargetJson() {
	libId, err := t.GetInt64(":lib")

	targets, err := models.NewTarget().GetTargetByLib(libId)

	t.Data["json"] = web.NewResponse(targets, err)
	t.ServeJSON()
}

// @Title CreateTarget
// @Description create target
// @Param	body		body 	models.Target	true		"body for target content"
// @Success 200 {int} models.Target.Id
// @Failure 403 body is empty
// @router / [post]
func (t *TargetController) AddTarget() {
	libId, err := t.GetInt64("id")
	if libId < 0 || err != nil {
		t.Abort("404")
	}

	if libId == 0 {
		libId = 1
	}

	defer func() {
		if err != nil {
			logs.Error(err)
		}
	}()

	name := t.GetString("name")
	if name == "" {
		t.JsonResult(4001, "目标名字不能为空")
	}

	age, err := t.GetInt64("age")
	if age > 140 || err != nil {
		t.JsonResult(4001, "请输入正确的年龄")
	}
	identity := t.GetString("identity")

	gender := t.GetString("gender")
	if gender != "male" && gender != "female" {
		t.JsonResult(4001, "请输入正确的性别")
	}

	nation := t.GetString("nation")       //民族
	host := t.GetString("host")           //籍贯
	message := t.GetString("description") //备注

	level, err := t.GetInt64("level")
	if level != 1 && level != 2 || err != nil {
		t.JsonResult(4001, "请输入正确的级别")
	}

	target := models.NewTarget()
	target.Name = name
	target.Identity = identity
	target.Gender = gender
	target.Level = level
	target.Age = age
	target.Nation = nation
	target.Host = host
	target.Message = message
	target.LibraryId = libId

	if _, err := target.LookUp(map[string]interface{}{"name": name}); err == nil {
		t.JsonResult(4001, "目标名字已存在")
	}

	// Photos
	photoStr := t.Input().Get("photo")
	var ps []models.Photo
	if err := json.Unmarshal([]byte(photoStr), &ps); err != nil {
		t.JsonResult(4001, err.Error())
	}

	if err := target.AddTarget(ps); err != nil {
		t.JsonResult(4001, err.Error())
	}

	t.JsonResult(0, "ok", target)
}

// @Title Update
// @Description update the target
// @Param	uid		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Target	true		"body for target content"
// @Success 200 {object} models.Target
// @Failure 403 :id is not int
// @router / [post]
func (t *TargetController) UpdateTarget() {
	t.TplName = "manager/edit_targets.tpl"
	id, err := t.GetInt64(":id")
	if id <= 0 || err != nil {
		t.Abort("404")
	}

	libId, err := t.GetInt64(":lib")
	if id <= 0 || err != nil {
		t.Abort("404")
	}

	defer func() {
		if err != nil {
			logs.Error(err)
		}
	}()

	target, err := models.NewTarget().LookUp(map[string]interface{}{"id": id})

	if t.Ctx.Input.IsPost() {
		name := t.GetString("name")
		if name == "" {
			t.JsonResult(4001, "目标名字不能为空")
		}

		age, err := t.GetInt64("age")
		if age > 140 || err != nil {
			t.JsonResult(4001, "请输入正确的年龄")
		}
		identity := t.GetString("identity")

		gender := t.GetString("gender")
		if gender != "male" && gender != "female" {
			t.JsonResult(4001, "请输入正确的性别")
		}

		nation := t.GetString("nation")       // 民族
		host := t.GetString("host")           //籍贯
		message := t.GetString("description") //备注

		level, err := t.GetInt64("level")
		if level != 1 && level != 2 || err != nil {
			t.JsonResult(4001, "请输入正确的级别")
		}

		target, _ := models.NewTarget().LookUp(map[string]interface{}{"id": id})
		target.Name = name
		target.Identity = identity
		target.Gender = gender
		target.Level = level
		target.Age = age
		target.Nation = nation
		target.Host = host
		target.Message = message
		target.LibraryId = libId

		// Photos
		photoStr := t.Input().Get("pictures")
		var ps []models.Photo
		if err := json.Unmarshal([]byte(photoStr), &ps); err != nil {
			t.JsonResult(4001, err.Error())
		}

		if err := target.UpdateTarget(ps); err != nil {
			t.JsonResult(4001, err.Error())
		}

		t.JsonResult(0, "ok", target)
	}

	fmt.Printf("???target:%+v\n\n", t.Ctx.Request)
	fmt.Printf("???target:%+v,%+v\n\n", libId, target)
	fmt.Printf("Pictures:%+v\n\n", target.Pictures)
	fmt.Printf("Pictures:%+v\n\n", target.Pictures[0])
	t.Data["Target"] = target
}

// @Title Delete
// @Description delete the target
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router / [post]
func (t *TargetController) DelTarget() {
	idStr := t.GetString("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {

		return
	}

	target, _ := models.NewTarget().LookUp(map[string]interface{}{"id": id})

	if err = target.DelTarget(); err != nil {
		logs.Error("删除目标 => ", err)
		t.JsonResult(4001, "删除失败", err)
	}
	t.JsonResult(0, "ok")
}

// Upload picture
// 1,Get Feature
// 2,copy to uploads/facial/
// 3,return fileName and fileUrl and feature
func (t *TargetController) UploadTarget() {
	t.Prepare()
	_, header, _ := t.GetFile("file")

	path, err := os.Create("/tmp/feat.jpg")
	defer path.Close()

	file, err := header.Open()
	defer file.Close()
	if err != nil {
	}

	if _, err := io.Copy(path, file); err != nil {
	}

	fid, err := models.GetFeature()
	if err != nil {
		data, err := models.CopyFileOnFacial("failed", header, fid)
		if err != nil {
		}
		t.JsonResult(6003, "failed", data)
	}

	data, err := models.CopyFileOnFacial("success", header, fid)
	if err != nil {
		t.JsonResult(6003, err.Error())
	}

	t.JsonResult(0, "ok", data)
}

func (t *TargetController) UploadTemp() {
	files, err := t.GetFiles("fileList")

	for i, _ := range files {
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
		}

		path, err := os.Create("uploads/facial/" + files[i].Filename)
		defer path.Close()
		if err != nil {
		}

		if _, err := io.Copy(path, file); err != nil {
		}
	}

	t.Data["Temp"] = "hoho"
	t.JsonResult(0, "ok", err)

}
