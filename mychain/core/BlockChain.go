package core

import (
	"fmt"
	"log"
)

type BlockChain struct {
	Blocks []*Block
}

//生成创世区块链
func NewBlockChain() *BlockChain {
	genesisBlock := GenerateGenesisBlock()
	blockChain := BlockChain{}
	blockChain.AppendBlock(genesisBlock)
	return &blockChain
}

//区块链运行
func (bc *BlockChain)SendData(data string){
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(preBlock,data)
	bc.AppendBlock(newBlock)
}

//区块链添加区块
func (bc *BlockChain) AppendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(newBlock, bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invalid block")
	}
}

//打印区块链信息
func (bc *BlockChain) Print() {
	for _, block := range bc.Blocks {
		fmt.Println("Index :",block.Index)
		fmt.Println("Prev.Hash :",block.PrevBlockHash)
		fmt.Println("Curr.Hash :",block.Hash)
		fmt.Println("BlockData :",block.Date)
		fmt.Println("Timestamp :",block.Timestamp)
		fmt.Println("-------------------------------------")
	}
}

//判断是否符合区块链协议
func isValid(newBlock, oldBlock *Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}
	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}
