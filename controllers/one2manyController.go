package controllers

import (
	_ "fmt"

	"github.com/lifei6671/mindoc/commands"
	"github.com/lifei6671/mindoc/conf"
	"github.com/lifei6671/mindoc/models"
	"github.com/lifei6671/mindoc/utils"
	"path/filepath"
)

// Operations about Uploads
type CompareController struct {
	ManagerController
}

// URLMapping ...
func (c *CompareController) URLMapping() {
	c.Mapping("Post", c.One2Many)
}

func (c *CompareController) One2ManySearch() {
	c.Prepare()
	c.TplName = "manager/one2manySearch.tpl"

	pageIndex, _ := c.GetInt("page", 1)

	attachList, totalCount, err := models.NewAttachment().FindToPager(pageIndex, conf.PageSize)

	if err != nil {
		c.Abort("500")
	}

	if totalCount > 0 {
		html := utils.GetPagerHtml(c.Ctx.Request.RequestURI, pageIndex, conf.PageSize, int(totalCount))

		c.Data["PageHtml"] = html
	} else {
		c.Data["PageHtml"] = ""
	}

	for _, item := range attachList {

		p := filepath.Join(commands.WorkingDirectory, item.FilePath)

		item.IsExist = utils.FileExists(p)

	}
	c.Data["Lists"] = attachList

}

// @Title Compare
// @Description one to many compare
// @Param   body        body    models.One2many  true        "The object content"
// @Success 200 {string}
// @Failure 403 body is empty
// @router / [post]
func (c *CompareController) One2Many() {
	threhold := c.GetString("score")
	lib := c.GetString("imgbaseid")
	file, _, _ := c.GetFile("file")

	res, _ := models.One2Many(threhold, lib, file)

	c.Data["json"] = &res

	c.ServeJSON()
}
