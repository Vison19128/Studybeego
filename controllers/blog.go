package controllers

import (
	"github.com/astaxie/beego"

	"Studybeego/models"
	"strings"
)

type BlogController struct {
	beego.Controller
}

func (c *BlogController) Get() {
	filePath := "blog/blog"
	c.TplName = "blog.html"
	reqUrl := c.Ctx.Request.URL.String()
	fullName := reqUrl[strings.LastIndex(reqUrl, "/")+1:]
	if qm := strings.Index(fullName, "?"); qm > -1 {
		fullName = fullName[:qm]
	}
	df := models.GetFile(filePath)
	if df == nil {
		c.Redirect("/blog", 302)
		return
	}

	c.Data["Title"] = df.Title
	c.Data["Data"] = string(df.Data)

}
