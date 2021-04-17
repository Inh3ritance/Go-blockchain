package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Data to store specified info
type Data struct {
	name  string
	_id   string
	value string
}

// Chain blocks together for DB
type Blockchain struct {
	blocks  []*Block
	counter int
}

// Block used inside chain
type Block struct {
	Timestamp     int64
	PrevBlockHash []byte
	Hash          []byte
	*Data
}

func Incriment(bc *Blockchain) {
	bc.counter++
}

func Decriment(bc *Blockchain) {
	bc.counter--
}

// Start chain if no block exists
func NewGenesisBlock() *Block {
	return NewBlock(Data{name: "Genesis Block", _id: "Generate ID", value: "N/A"}, []byte{})
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, []byte(b.Data._id), timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data Data, prevBlockHash []byte) *Block {
	data._id = strconv.Itoa(rand.Intn(100))
	block := &Block{time.Now().Unix(), prevBlockHash, nil, &data}
	block.SetHash()
	return block
}

func (bc *Blockchain) AddBlock(data Data) {
	if data.value == "up" {
		Incriment(bc)
	} else if data.value == "down" {
		Decriment(bc)
	}

	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}, 0}
}

func main() {
	rand.Seed(20)
	bc := NewBlockchain()

	bc.AddBlock(Data{name: "Josh", value: "up"})
	bc.AddBlock(Data{name: "Drake", value: "up"})

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

	fmt.Printf("counter: %x\n", bc.counter)
}
