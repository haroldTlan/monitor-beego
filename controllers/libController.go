package controllers

import (
	"encoding/json"
	_ "fmt"
	"html/template"
	"strconv"
	"strings"

	"github.com/lifei6671/mindoc/models"
)

// 希望接手的下一位，不要打我，都是前端的锅，我也是很累
// Operations about Librarys
type LibController struct {
	ManagerController
}

// URLMapping ...
func (l *LibController) URLMapping() {
	l.Mapping("GetAll", l.Librarys)
	l.Mapping("Post", l.AddLibrary)
	l.Mapping("Post", l.UpdateLibrary)
	l.Mapping("Post", l.DelLibrary)
}

// @Title GetAll
// @Description get all librarys
// @Success 200 {object} models.Library
// @router / [get]
func (l *LibController) Librarys() {
	l.Prepare()
	l.TplName = "manager/librarys.tpl"

	libs, err := models.GetAllLibrarys()

	b, err := json.Marshal(libs)

	if err != nil {
		l.Data["Results"] = template.JS("[]")
	} else {
		l.Data["Results"] = template.JS(string(b))
	}
	l.Data["Libs"] = libs
}

// @Title Create
// @Description create library
// @Param	body		body 	models.Library	true		"The object content"
// @Success 200 {string} models.Library.Id
// @Failure 403 body is empty
// @router / [post]
func (l *LibController) AddLibrary() {
	name := strings.TrimSpace(l.GetString("name"))
	if name == "" {
		l.JsonResult(4000, "目标库名字不能为空")
	}

	role, err := l.GetInt64("role")
	if err != nil {
		l.JsonResult(4000, "请输入正确的目标库种类")
	}

	message := strings.TrimSpace(l.GetString("description"))
	if len(message) >= 140 {
		l.JsonResult(4000, "目标库备注长度不能超过140字")
	}

	library := models.NewLibrary()
	if err := library.CheckLibraryByName(name); err != nil {
		l.JsonResult(4000, err.Error())
	}

	if role != 0 && role != 1 && role != 2 {
		role = 1
	}

	library.Name = name
	library.RoleId = role
	library.Message = message

	if err := library.AddLibrary(); err != nil {
		l.JsonResult(4000, "添加目标库失败", err)
	}

	l.JsonResult(0, "ok", library)
}

// @Title Update
// @Description update the library
// @Param	name		path 	string	true		"The name you want to update"
// @Param	role		path 	string	true		"The role you want to update"
// @Param	message		path 	string	true		"The message you want to update"
// @Success 200 {object} models.UpdateLib
// @Failure 403 :objectId is empty
// @router /:id [put]
func (l *LibController) UpdateLibrary() {
	l.Prepare()
	l.TplName = "manager/edit_librarys.tpl"

	id, err := l.GetInt64(":id")
	if id <= 0 || err != nil {
		l.Abort("404")
	}

	library, err := models.NewLibrary().CheckLibraryById(id)

	if l.Ctx.Input.IsPost() {
		name := strings.TrimSpace(l.GetString("name"))
		role, err := l.GetInt64("role")
		if err != nil {
			l.JsonResult(4000, "请输入正确的目标库种类")
		}
		message := strings.TrimSpace(l.GetString("description"))

		if err := library.CheckLibraryByName(name); err != nil {
			l.JsonResult(4000, err.Error())
		}

		if role != 0 && role != 1 && role != 2 {
			role = 1
		}

		if len(message) > 100 {
			l.JsonResult(4000, "备注超过100字")
		}

		library.Name = name
		library.RoleId = role
		library.Message = message

		if err = library.UpdateLibrary(); err != nil {
			l.JsonResult(4000, "保存失败", err)
		}

		l.JsonResult(0, "ok")
	}
	l.Data["Lib"] = library
}

// @Title Delete
// @Description delete the library
// @Param	library		path 	string	true		"The library id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Id is empty
// @router /:id [delete]
func (l *LibController) DelLibrary() {
	id, _ := l.GetInt64("id")

	library, err := models.NewLibrary().CheckLibraryById(id)
	if err != nil {
		l.JsonResult(4000, "不存在的目标库ID: "+strconv.FormatInt(id, 10))
	}

	if err = library.DelLibrary(); err != nil {
		l.JsonResult(4000, "删除失败", err)
	}
	l.JsonResult(0, "ok")
}
