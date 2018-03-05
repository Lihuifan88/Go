package main

import (
	_ "study/day20_beego/WEB/routers" //_ 表示只导入默认方法

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
