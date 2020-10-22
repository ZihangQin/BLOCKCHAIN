package controllers

import (
	"DataCertProject/models"
	"DataCertProject/util"
	//"crypto/sha256"
	//"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"path"
	"time"
)

type UploadController struct {
	beego.Controller
}

func (u *UploadController) Get() {
	phone := u.GetString("phone")
	u.TplName = "attestation.html"
	u.Data["Phone"] = phone
}

func (u *UploadController) Post() {
	//fileTitle := u.Ctx.Request.PostFormValue("upload_title")
	////文件
	//file, header, err :=u.GetFile("upload_file")
	//if err != nil {
	//	u.Ctx.WriteString("抱歉，用户文件解析失败，请重试")
	//	return
	//}
	//fmt.Println("自定义的文件标题:",fileTitle)
	//fmt.Println("文件名称:",header.Filename)
	//fmt.Println("文件的大小:",header.Size)
	//
	//fmt.Println(file)
	//u.Ctx.WriteString("解析到上传文件，文件名是："+header.Filename)
	////2、将文件保存在本地的一个目录中
	////文件全路径： 路径 + 文件名 + "." + 扩展名
	////要的文件的路径
	//err = u.SaveToFile("upload_file", "./static/img/"+header.Filename+".jpg")
	//if err != nil {
	//	fmt.Println(err.Error())
	//	u.Ctx.WriteString("上传失败")
	//} else {
	//	u.Ctx.WriteString("上传成功")
	//}
	//
	//defer file.Close()

	phone := u.Ctx.Request.PostFormValue("phone")
	//1、获取文件
	f, h, err := u.GetFile("upload_file")
	filesTitle := u.Ctx.Request.PostFormValue("upload_title")
	if err != nil {
		u.Ctx.WriteString("获取文件失败，请从新上传。")
		return
	}
	//2、上传文件，判断文件的打小以及格式
	if h.Size > 10000000 {
		u.Ctx.WriteString("单次上传文件太大，请压缩后再进行上传！")
		fmt.Println(h.Size)
	} else {
		ext := path.Ext(h.Filename)
		//fmt.Println("自定义的文件标题:", filesTitle)
		//fmt.Println("文件名称:", h.Filename)
		//fmt.Println("文件的大小:", h.Size)
		//fmt.Println(f)
		var AllowExt map[string]bool = map[string]bool{
			//增加规则，只能上传文件类型是图片的文件
			".jpg":  true,
			".png":  true,
			".jpeg": true,
			".psd":  true,
		}
		if _, ok := AllowExt[ext]; !ok {
			u.Ctx.WriteString("后缀名不符合要求")
			return
		}
		//3、把文件名进行加密处理
		hash, err := util.MD5HashReader(f)
		uploadDir := "./static/img/save/" + h.Filename
		err = u.SaveToFile("upload_file", uploadDir)
		if err != nil {
			fmt.Println(err.Error())
			u.Ctx.WriteString("上传失败")
		}
		//else {
		//u.Ctx.WriteString("解析到上传文件：" + h.Filename)
		//}
		defer f.Close()

		//now := time.Now().Unix()
		//tm := time.Unix(now,0)
		//time := tm.Format("2006/01/02 15:04:05")
		//times,err := strconv.FormatInt(tm, 10)


		record := models.UploadFile{}
		record.FileName = h.Filename
		record.FileSize = h.Size
		record.FileTitle = filesTitle
		record.Time = time.Now().Unix()
		record.FileHash = hash
		record.User_phone = phone
		//t2 := time.Unix(times,0).Format("2006/01/02 15:04:05")

		_, err = record.SaveUploadRecord()
		if err != nil {
			fmt.Println(err.Error())
			u.Ctx.WriteString("文件信息保存错误")
			return
		}
	//	查询数据库中的数据
		records, err := models.QueryRecordByPhone(phone)

		if err != nil {
			fmt.Println(err.Error())
			u.Ctx.WriteString("获取数据出现错误错误")
			return
		}
		//record.Time = time.Format("2006/01/02 15:04:05")
		u.Data["Phone"] = phone
		u.Data["Records"] = records
		u.TplName = "upload_list.html"
		//u.Ctx.WriteString("获取文件")
	}
}
