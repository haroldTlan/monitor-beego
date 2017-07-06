package controllers

import (
	_ "encoding/json"
	_ "fmt"
	"strconv"

	"github.com/astaxie/beego"
	"last/controllers/web"
	"last/models"
)

// Librarys about object
type LibController struct {
	beego.Controller
}

// URLMapping ...
func (l *LibController) URLMapping() {
	l.Mapping("GetAll", l.GetAll)
	l.Mapping("Post", l.Post)
	l.Mapping("Put", l.Update)
	l.Mapping("Delete", l.Delete)
}

// @Title GetAll
// @Description get all librarys
// @Success 200 {object} models.Library
// @Failure 403 :libId is empty
// @router / [get]
func (l *LibController) GetAll() {
	res, err := models.GetAllLibrarys()
	l.Data["json"] = web.NewResponse(res, err)
	l.ServeJSON()
}

// @Title Create
// @Description create library
// @Param	body		body 	models.Library	true		"The object content"
// @Success 200 {string} models.Library.Id
// @Failure 403 body is empty
// @router / [post]
func (l *LibController) Post() {
	name := l.GetString("name")
	role := l.GetString("role")
	message := l.GetString("message")

	err := models.AddLibrary(name, role, message)
	l.Data["json"] = web.NewResponse("", err)
	l.ServeJSON()
}

// @Title Update
// @Description update the library
// @Param	name		path 	string	true		"The name you want to update"
// @Param	role		path 	string	true		"The role you want to update"
// @Param	message		path 	string	true		"The message you want to update"
// @Success 200 {object} models.UpdateLib
// @Failure 403 :objectId is empty
// @router /:id [put]
func (l *LibController) Update() {
	var err error
	defer func() {
		l.Data["json"] = web.NewResponse("update success!", err)
		l.ServeJSON()
	}()

	idStr := l.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return
	}

	name := l.GetString("name")
	role := l.GetString("role")
	message := l.GetString("message")

	err = models.UpdateLibrary(name, role, message, id)
}

// @Title Delete
// @Description delete the library
// @Param	library		path 	string	true		"The library id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Id is empty
// @router /:id [delete]
func (l *LibController) Delete() {
	var err error
	defer func() {
		l.Data["json"] = web.NewResponse("", err)
		l.ServeJSON()
	}()

	idStr := l.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return
	}
	err = models.DelLibrary(id)
}
