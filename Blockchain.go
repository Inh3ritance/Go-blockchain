package main

import (
	"fmt"
	"math/rand"
)

// Global key, bad practice only for example
var KeyAES = []byte("12345678901234567890123456789012")

func main() {
	rand.Seed(20)
	bc := NewBlockchain()

	bc.AddBlock(Data{name: "Josh", value: "up"})
	bc.AddBlock(Data{name: "Drake", value: "up"})
	bc.AddBlock(Data{name: "Robin", value: "down"})

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

	fmt.Printf("counter results: %x\n", bc.counter)

	fmt.Printf("Verify some info")
	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		if block.name == "Genesis Block" {
			fmt.Printf("Data-name: %s\n", block.name)
		} else {
			fmt.Printf("Data-name: %s\n", RemovePadding(DecryptAES(KeyAES, block.name)))
		}
		fmt.Printf("Data-value: %s\n", block.value) // dont need block.Data.name can skip to block.name
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
