package controllers

import (
	_ "fmt"

	"github.com/astaxie/beego"
	"github.com/lifei6671/mindoc/models"
)

// Operations about Librarys
type TempController struct {
	beego.Controller
}

// URLMapping ...
func (l *TempController) URLMapping() {
	l.Mapping("GetAll", l.Librarys)
}

type Results struct {
	Name   string `json:"result"`
	Result `json:"data"`
}

type Result struct {
	Total int              `json:"total"`
	Data  []models.Library `json:"data"`
}

// @Title GetAll
// @Description get all librarys
// @Success 200 {object} models.Library
// @router / [get]
func (l *TempController) Librarys() {
	//str:=`{createtime:"2016-07-20 16:31:26",id:2,name:"测试图片库",number:1591,remark:""]}`
	libs, _ := models.GetAllLibrarys()

	l.Data["json"] = &Results{Result: Result{Total: len(libs), Data: libs}, Name: "success"}
	l.ServeJSON()
}
