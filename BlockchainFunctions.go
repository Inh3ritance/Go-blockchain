package main

import (
	"bytes"
	"crypto/sha256"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Add to Blockchain counter
func Incriment(bc *Blockchain) {
	bc.counter++
}

// Subtract from Blockchain counter
func Decriment(bc *Blockchain) {
	bc.counter--
}

// Start new Ledger
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}, 0}
}

// Start a new ledger Block
func NewGenesisBlock() *Block {
	return NewBlock(Data{name: "Genesis Block", _id: "Generate ID", value: "N/A"}, []byte{})
}

// after setting block hash data
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, []byte(b.Data._id), timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// Assign new ID to Block and pass Data to Block
func NewBlock(data Data, prevBlockHash []byte) *Block {
	data._id = strconv.Itoa(rand.Intn(100))
	block := &Block{time.Now().Unix(), prevBlockHash, nil, &data}
	block.SetHash()
	return block
}

// add padding 16bytes for cbc
func Padding(str string) string {
	for len(str) < 16 {
		str += "-"
	}
	return str
}

// remove Padding
func RemovePadding(str string) string {
	for str[len(str)-1:] == "-" {
		str = strings.TrimSuffix(str, "-")
	}
	return str
}

// Create a new block and evaluate whether to incriment or decriment counter
// Dynamic set func exclusive for bc.Addblock
func (bc *Blockchain) AddBlock(data Data) {
	data.name = Padding(data.name)
	data.name = EncryptAES(KeyAES, data.name)
	if data.value == "up" {
		Incriment(bc)
	} else if data.value == "down" {
		Decriment(bc)
	}
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
