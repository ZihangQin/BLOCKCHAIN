package models

import (
	"DataCertProject/db_mysql"
	"DataCertProject/util"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type User struct {
	Id int          `form:"id"`
	Phone string    `form:"phone"`
	Password string `form:"password"`
	Time string    `form:"time"`
}
//保存到数据库信息
func (u User) SeverUser() (int64,error) {
	//1、密码脱敏处理
	//hashSHA256:=sha256.New()
	//hashSHA256.Write([]byte(u.Password))
	//bytes := hashSHA256.Sum(nil)
	u.Password = util.MD5HashString(u.Password)
	//上传创建用户的时间
	var timeSilce []byte
	timeSilce = []byte(time.Now().String())
	time := string(timeSilce[:19])
	u.Time = time
	//2、执行数据库操作
	row,err := db_mysql.DB.Exec("insert into user (phone,password,time) values(?,?,?)",u.Phone,u.Password,u.Time)
	if err!=nil {
		return -1,err
	}
	id,err := row.RowsAffected()
	if err != nil {
		return -1,err
	}
	return id,nil
}



//查询数据库信息
func (u User) QueryUser() (*User,error) {
	hashSHA256:=sha256.New()
	hashSHA256.Write([]byte(u.Password))
	bytes := hashSHA256.Sum(nil)
	u.Password = hex.EncodeToString(bytes)
	row := db_mysql.DB.QueryRow("select phone from user where phone = ? and password = ?",u.Phone,u.Password)
	err := row.Scan(&u.Phone)
	if err != nil {
		return nil,err
	}
	return &u,nil
}