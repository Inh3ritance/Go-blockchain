package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Global key, bad practice only for example
var KeyAES = []byte("12345678901234567890123456789012")

func main() {

	// Create 2 WaitGroups for 2 GoRoutines
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		fmt.Println("Los Angeles")
		rand.Seed(20)
		bc := NewBlockchain()

		bc.AddBlock(Data{name: "Josh", value: "up"})
		bc.AddBlock(Data{name: "Drake", value: "up"})
		bc.AddBlock(Data{name: "Robin", value: "down"})
		bc.AddBlock(Data{name: "Mark", value: "down"})
		bc.AddBlock(Data{name: "Sharron", value: "up"})
		bc.AddBlock(Data{name: "Terry", value: "up"})

		for _, block := range bc.blocks {
			fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
			fmt.Printf("Data: %s\n", block.Data)
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Println()
		}

		fmt.Printf("counter results LOS ANGELES: %x\n", bc.counter)

		fmt.Printf("Verify some info")
		for _, block := range bc.blocks {
			fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
			if block.name == "Genesis Block" {
				fmt.Printf("Data-name: %s\n", block.name)
			} else {
				fmt.Printf("Data-name: %s\n", RemovePadding(DecryptAES(KeyAES, block.name)))
			}
			fmt.Printf("Data-value: %s\n", block.value)
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Println()
		}
		wg.Done()
	}()

	go func() {
		fmt.Println("New York")
		rand.Seed(21)
		bc := NewBlockchain()

		bc.AddBlock(Data{name: "Kirk", value: "down"})
		bc.AddBlock(Data{name: "Daniel", value: "down"})
		bc.AddBlock(Data{name: "Simon", value: "down"})
		bc.AddBlock(Data{name: "GeekBench", value: "down"})
		bc.AddBlock(Data{name: "random", value: "up"})
		bc.AddBlock(Data{name: "hacker", value: "down"})

		for _, block := range bc.blocks {
			fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
			fmt.Printf("Data: %s\n", block.Data)
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Println()
		}

		fmt.Printf("counter results NEW YORK: %x\n", bc.counter)

		fmt.Printf("Verify some info")
		for _, block := range bc.blocks {
			fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
			if block.name == "Genesis Block" {
				fmt.Printf("Data-name: %s\n", block.name)
			} else {
				fmt.Printf("Data-name: %s\n", RemovePadding(DecryptAES(KeyAES, block.name)))
			}
			fmt.Printf("Data-value: %s\n", block.value)
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Println()
		}
		wg.Done()
	}()

	wg.Wait()
}
