package controllers

import (
	"github.com/astaxie/beego"
	"last/controllers/web"
	"last/models"
)

// Operations about Targets
type TargetController struct {
	beego.Controller
}

// URLMapping ...
func (t *TargetController) URLMapping() {
	t.Mapping("GetAll", t.GetAll)
	t.Mapping("Post", t.Post)
	/*

		l.Mapping("Put", l.Update)
		l.Mapping("Delete", l.Delete)
	*/
}

// @Title GetAll
// @Description get all Targets
// @Success 200 {object} models.Target
// @router / [get]
func (t *TargetController) GetAll() {
	targets, err := models.GetAllTargets()
	t.Data["json"] = web.NewResponse(targets, err)
	t.ServeJSON()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (t *TargetController) Post() {
	name := t.GetString("name")
	identity := t.GetString("identity")
	sex := t.GetString("sex")
	nation := t.GetString("nation")
	host := t.GetString("host")
	message := t.GetString("message")
	library, _ := t.GetInt64("library")
	level, _ := t.GetInt64("level")
	age, _ := t.GetInt64("age")
	file, _, _ := t.GetFile("file")

	err := models.AddTarget(name, identity, sex, nation, host, message, level, age, library, file)

	t.Data["json"] = web.NewResponse("", err)
	t.ServeJSON()
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
