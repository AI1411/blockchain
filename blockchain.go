package main

import (
	"fmt"
	"time"
)

//blockを定義
type Block struct {
	nonce int
	previousHash string
	timestamp int64
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

func main()  {
	b := NewBlock(0, "init hash")
	b.Print()
}