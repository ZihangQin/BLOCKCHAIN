package routers

import (
	"DataCertProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//设置用户注册页面接口
    beego.Router("/", &controllers.MainController{})
    //设置用户注册请求接口
    beego.Router("/user_register", &controllers.RegisterController{})
    //设置用户直接登录页面接口
    beego.Router("/login.html",&controllers.LoginController{})
	//设置存证页面接口
    beego.Router("/attestation",&controllers.LoginController{})
	//设置存证接口
    beego.Router("/attestation_up",&controllers.UploadController{})
    //添加新增页面
    beego.Router("/upload_file.html",&controllers.UploadController{})
}
