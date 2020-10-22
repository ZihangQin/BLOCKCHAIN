package controllers

import (
	"DataCertProject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post() {
	//1、解析请求数据
	var user models.User
	//errTmp , _:= template.ParseFiles("./views/error.html")
	err := r.ParseForm(&user)
	//fmt.Println(user)
	if err != nil {
		//errTmp.Execute(log.Writer(),"解析数据出错")
		r.Ctx.WriteString("解析数据出错！")
		return
	}
	//2、保存用户信息到数据库
	_, err = user.SeverUser()
	//3、返回前端结果
	//注册失败
	if err != nil {
		//errTmp.Execute(log.Writer(),"用户注册失败")
		r.Ctx.WriteString("用户注册失败")
		fmt.Println(err.Error())
		return
	}
	//注册成功
	r.TplName = "login.html"
}