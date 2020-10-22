package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

/**
 * 对一个字符串进行MD5哈希计算,并返回hash值
 */
func MD5HashString(data string) string {
	md5Hash := sha256.New()
	md5Hash.Write([]byte(data))
	psswordBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(psswordBytes)
}

/**
 * 对一个io操作的reader（通常为文件)进行数据读取，并计算hash，返回md5哈希值
 */
func MD5HashReader(reader io.Reader) (string, error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	md5Hash := sha256.New()
	md5Hash.Write(bytes)
	hashBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}

/**
 * 对数据进行hash计算
 */
func SHA256Hash(data []byte) ([]byte) {
	//、对数据进行sha256
	sha256Hash := sha256.New()
	sha256Hash.Write(data)
	return sha256Hash.Sum(nil)
}
