package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type UserInfo struct {
	Id         int64
	Username   string
	Departname string
}

type TestModelController struct {
	beego.Controller
}

func (c *TestModelController) Get() {
	orm.Debug = true

	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/testinfo?charset=utf8", 30)
	orm.RegisterModel(new(UserInfo))
	o := orm.NewOrm()
	//	user := UserInfo{Username: "李四", Departname: "开发部"}
	//插入
	//	id, err := o.Insert(&user)
	//修改
	//	user.Id = 2
	//	user.Username = "李四改"
	//	id, err := o.Update(&user)
	//	c.Ctx.WriteString(fmt.Sprintf("id: %d, err: %v", id, err))
	//读
	//	user.Id = 2
	//	o.Read(&user)
	//	c.Ctx.WriteString(fmt.Sprintf("user info: %v", user))

	//下面是原生读取
	//	var maps []orm.Params
	//	o.Raw("select * from user_info").Values(&maps)
	//	var users []UserInfo
	//	o.Raw("select * from user_info").QueryRows(&users)
	//	//	for _, v := range maps {
	//	c.Ctx.WriteString(fmt.Sprintf("user info: %v", users))
	//	}

	//通过QueryBuilder查询
	var users []UserInfo
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("user_info").Where("username=?").And("id=?").Limit(1)
	sql := qb.String()
	o.Raw(sql, "张三", 1).QueryRows(&users)
	c.Ctx.WriteString(fmt.Sprintf("user info: %v", users))
}
