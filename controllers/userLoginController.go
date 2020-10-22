package controllers

import (
	"DataCertProject/models"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}
//
func (l *LoginController) Get() {
	//l.TplName = "error.html"
		//需要设置页面
	l.TplName = "login.html"

}

func (l *LoginController) Post()  {
	var user models.User
	//errTmp , _:= template.ParseFiles("./views/error.html")
	err := l.ParseForm(&user)
	if err != nil {
		//errTmp.Execute(nil,"用户信息解析失败",)
		l.Ctx.WriteString("用户信息解析失败" )
		return
	}
	//查询用户信息
	u ,err :=  user.QueryUser()
	if err != nil {
		//fmt.Println(err.Error())
		//errTmp.Execute(io.Writer(nil),"用户登录失败")
		l.Ctx.WriteString("用户登录失败")
		 return
	}
	//登录成功跳转页面
	l.Data["Phone"] = u.Phone
	l.TplName = "attestation.html"
}