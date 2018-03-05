package controllers

import (
	"github.com/astaxie/beego"
)

type TestSession struct {
	beego.Controller
}

func (c *TestSession) Get() {
	userName := c.GetSession("userName")
	if nameStr, ok := userName.(string); ok && nameStr != "" {
		c.Ctx.WriteString("用户" + userName.(string) + "获取Session登陆成功")
	} else {
		c.Ctx.WriteString(`<html><form action="http://127.0.0.1:8080/test_request" method="post">
  			<p>用户名<input type="text" name="Username" /></p>
  			<p>密码<input type="text" name="Password" /></p>
  			<input type="submit" value="登陆"/>
		</form></html>`)
	}
}

func (c *TestSession) Post() {
	user := User{}
	if err := c.ParseForm(&user); err != nil {
		//处理异常
	}
	c.Ctx.WriteString("用户" + user.Username + "登陆成功")
}
