package models

import (
	"DataCertProject/db_mysql"
	"fmt"
	"time"
)

type UploadFile struct {
	Id         int    `form:"id"`
	FileName   string `form:"fileName"`
	FileSize   int64  `form:"fileSize"`
	FileHash   string `form:"fileHash"`
	FileTitle  string `form:"fileTitle"`
	Time       int64  `form:"time"`
	Size	   float64
	TimeFormat string
	User_phone string `form:"user_phone"`
}

func (u UploadFile) SaveUploadRecord() (int64, error) {
	rs, err := db_mysql.DB.Exec("insert into upload(fileName,"+
		"fileSize,fileHash,fileTitle,time,user_phone) values(?,?,?,?,?,?)",
		u.FileName, u.FileSize, u.FileHash, u.FileTitle, u.Time, u.User_phone)
	if err != nil {
		return -1, err
	}
	id, err := rs.RowsAffected()
	if err != nil {
		return id, nil
	}
	return id, nil
}

func QueryRecordByPhone(phone string) ([]UploadFile, error) {
	rs, err := db_mysql.DB.Query("select id,fileName,fileSize,fileHash, "+
		"fileTitle,time,user_phone from upload where user_phone = ?", phone)
	if err != nil {
		return nil, err
	}
	records := make([]UploadFile, 0)
	for rs.Next() {
		var reord UploadFile
		err := rs.Scan(&reord.Id, &reord.FileName, &reord.FileSize, &reord.FileHash,
			&reord.FileTitle, &reord.Time, &reord.User_phone)
		if err != nil {
			return nil, err
		}

		b := float64(reord.FileSize)/float64(1024*1024)
		fmt.Printf("%.2fMB\n",b)
		reord.Size = b

		a := time.Unix(reord.Time, 0).Format("2006/01/02 15:04:05")
		//fmt.Println(a)
		reord.TimeFormat = a
		records = append(records, reord)

		//tm := time.Unix(reord.Time,0)
		//tm := strconv.FormatInt(reord.Time,10)
		//tm.Format("2006/01/02 15:04:05")
	}
	//now := time.Now().Unix()
	//tm := time.Unix(now,0)
	//time := now.Format("2006/01/02 15:04:05")
	//times,err := strconv.FormatInt(tm, 10)

	return records, nil
}
