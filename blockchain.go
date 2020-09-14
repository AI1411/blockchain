package main

import (
	"fmt"
	"strings"
	"time"
)

//blockを定義
type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	//スライス型
	transactions []string
}

//メソッドを定義
func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp     %d\n", b.timestamp)
	fmt.Printf("nonce         %d\n", b.nonce)
	fmt.Printf("transactions  %s\n", b.transactions)
	fmt.Printf("previous Hash %s\n", b.previousHash)
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

//blockchainを生成し、初期値をいれる
func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

//BlockchainのプールにBlockを追加していく処理
func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

//NewBlockchain表示用
func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func main() {
	blockchain := NewBlockchain()
	blockchain.Print()
	blockchain.CreateBlock(5, "hash 1")
	blockchain.Print()
	blockchain.CreateBlock(2, "hash 2")
	blockchain.Print()
}
