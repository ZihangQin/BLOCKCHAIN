package block_chain

import (
	"bytes"
	"encoding/gob"
	"time"
)

type Block struct {
	//上一块的hash
	PreHash []byte
	//当前区块的hash
	HashCode []byte
	//时间戳
	TimeStamp int64
	//当前网络的难读系数,控制hash有几个前导0的
	//Diff int
	//存储交易信息
	Data []byte
	//区块高度
	Index int64
	//随机值
	Nonce int64
	//版本号
	Version string
}

//生成创世块
func CreateGenesisBlock() Block {
	block := NewBlock(0, []byte{}, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	return block
}




//创建区块块
func NewBlock(h int64, data []byte,p []byte) Block {

	var block Block
	block.TimeStamp = time.Now().Unix()
	block.Data = data
	block.Index = 1
	block.Version = "001"
	block.PreHash = p

	pow := NewPow(block)
	blockchain,nonce := pow.Run()

	//问题分析：
	//① util.SHA256Hash要求一个[]byte参数
	//② block是一个自定义结构体, 与①类型不匹配

	//解决思路：将block结构体转换为[]byte类型数据

	block.Nonce = nonce
	block.HashCode = blockchain

	return block
}



//序列化区块
func (bk Block) Serialize() ([]byte, error) {
	buff := new(bytes.Buffer)
	err := gob.NewEncoder(buff).Encode(bk)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

//反序列化区块
func DeSerialize(data []byte) (*Block, error) {
	var block Block
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

