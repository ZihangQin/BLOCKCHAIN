package main

import (
	"DataCertProject/block_chain"
	"DataCertProject/db_mysql"
	_ "DataCertProject/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/boltdb/bolt"
)

var BUCKET_NAME = "blocks"

func main() {

	block := block_chain.CreateGenesisBlock()
	fmt.Println(block)
	fmt.Printf("区块Hash值：%x\n", block.HashCode)
	fmt.Printf("区块的nonce值：%d\n", block.Nonce)
	fmt.Printf("区块的时间戳值：%d\n", block.TimeStamp)

	db,err := bolt.Open("chain.db",0600,nil)
	if err != nil {
		//fmt.Println(err.Error())
		panic(err.Error())
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		var b *bolt.Bucket

		b = tx.Bucket([]byte(BUCKET_NAME))
		if b == nil {
			b, err  = tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				return err
			}
		}
		v := b.Get([]byte("lasthash"))
		blockHash, err:= block.Serialize()
		//序列化创世块
		if err != nil {
			return err
		}
		if v == nil {
			b.Put(block.HashCode,blockHash)

			b.Put([]byte("lasthash"),blockHash)
		}
		err = b.Put([]byte(""),[]byte(""))
		return err
	})
	return

	//1、开启数据库连接
	db_mysql.ConnectDB()
	//加载静态文件
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
	beego.Run()
}
