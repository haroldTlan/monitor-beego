package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"

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
	t.Prepare()
	t.TplName = "manager/targets.tpl"

	libId, _ := t.GetInt64(":lib")

	lib, _ := models.NewLibrary().CheckLibraryById(libId)

	targets, err := models.NewTarget().GetTargetByLib(libId)
	fmt.Printf("%+v", targets)
	b, err := json.Marshal(targets)

	if err != nil {
		t.Data["Results"] = template.JS("[]")
	} else {
		t.Data["Results"] = template.JS(string(b))
	}

	t.Data["Targets"] = targets
	t.Data["Lib"] = lib
}

// @Title CreateTarget
// @Description create target
// @Param	body		body 	models.Target	true		"body for target content"
// @Success 200 {int} models.Target.Id
// @Failure 403 body is empty
// @router / [post]
func (t *TargetController) AddTarget() {
	target := models.NewTarget()

	target.Name = t.GetString("name")
	target.Identity = t.GetString("identity")
	target.Sex = t.GetString("sex")
	target.Nation = t.GetString("nation")
	target.Host = t.GetString("host")
	target.Message = t.GetString("description")
	target.LibraryId, _ = t.GetInt64("library")
	target.Level, _ = t.GetInt64("level")
	target.Age, _ = t.GetInt64("age")
	//file, _, _ := t.GetFile("file")
	fmt.Printf("%+v", target)
	t.JsonResult(0, "ok", target)
}

// @Title Update
// @Description update the target
// @Param	uid		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Target	true		"body for target content"
// @Success 200 {object} models.Target
// @Failure 403 :id is not int
// @router /:id [post]
func (u *TargetController) UpdateTarget() {
}

// @Title Delete
// @Description delete the target
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [post]
func (u *TargetController) DelTarget() {
}

/*

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models.Login(username, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

*/
