package main

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
