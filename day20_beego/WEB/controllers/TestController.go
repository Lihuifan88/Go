package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

type User struct {
	Username string
	Password string
}

func (c *TestController) Get() {

	userName := c.Ctx.GetCookie("userName")
	if userName != "" {
		c.Ctx.WriteString("用户" + userName + "登陆成功")
	} else {
		c.Ctx.WriteString(`<html><form action="http://127.0.0.1:8080/test_request" method="post">
  			<p>用户名<input type="text" name="Username" /></p>
  			<p>密码<input type="text" name="Password" /></p>
  			<input type="submit" value="登陆"/>
		</form></html>`)
	}
}

func (c *TestController) Post() {
	user := User{}
	if err := c.ParseForm(&user); err != nil {
		//处理异常
	}
	c.Ctx.SetCookie("userName", user.Username, 100, "/")

	//Session
	c.SetSession("userName", user.Username)
	c.Ctx.WriteString("用户" + user.Username + "登陆成功")
}
