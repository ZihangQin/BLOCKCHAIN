package block_chain

import (
	"DataCertProject/util"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	Target	*big.Int
	Block Block
}

const diff  = 16



//设置系统给定的哈希值
func NewPow(block Block) ProofOfWork {
	t := big.NewInt(1)
	t.Lsh(t,255-diff)

	pow := ProofOfWork{
		Target: t,
		Block:block,
			}
			return pow
}

func (p ProofOfWork) Run() ([]byte , int64) {
	bigBlock := new(big.Int)
	var nonce int64
	var block256Hash []byte
	for  {
		var block Block
		block = p.Block

		heightBytes,_ := util.IntToBytes(block.Index)
		timeBytes,_ := util.IntToBytes(block.TimeStamp)
		versionBytes := util.StringToBytes(block.Version)
		nonceBytes,_ := util.IntToBytes(nonce)

		blockByets := bytes.Join([][]byte{
			heightBytes,
			timeBytes,
			block.Data,
			block.PreHash,
			versionBytes,
			nonceBytes,
		},[]byte{})

		sha256Hash := sha256.New()
		sha256Hash.Write(blockByets)

		block256Hash = sha256Hash.Sum(nil)

		fmt.Printf("挖矿正在进行，当前的nonce值为：%d\n",nonce)
		bigBlock = bigBlock.SetBytes(block256Hash)
		fmt.Printf("预期值：%x\n",p.Target)
		fmt.Printf("hash值：%x\n",bigBlock)


		if p.Target.Cmp(bigBlock) == 1{
			fmt.Println("挖矿成功")
			break

		}
			nonce++

	}
	return block256Hash,nonce
}
