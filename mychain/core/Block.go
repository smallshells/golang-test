package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// 定义区块
type Block struct {
	//区块头
	Index         int64  //区块编号
	Timestamp     int64  //区块时间戳
	PrevBlockHash string //上一个区块哈希值
	Hash          string //当前区块哈希值

	//区块体
	Date string //区块数据
}

// 哈希计算函数
func CalculateHash(b *Block) string {
	blockDate := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Date
	hashInBytes := sha256.Sum256([]byte(blockDate))
	hashInStr := hex.EncodeToString(hashInBytes[:])
	return hashInStr
}

//生成一个区块
func GenerateNewBlock(preBlock *Block, data string) *Block {
	newBlock := &Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Date = data
	newBlock.Hash = CalculateHash(newBlock)
	return newBlock
}

//生成创世区块
func GenerateGenesisBlock() *Block {
	preBlock := &Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock, "Genesis Block")
}
