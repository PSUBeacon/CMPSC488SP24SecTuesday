package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

// Block represents a single block in the blockchain.
type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

// Blockchain is a simple blockchain structure.
type Blockchain struct {
	Chain []Block
}

// CalculateHash calculates the SHA-256 hash of a block.
func CalculateHash(block Block) string {
	data := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp, block.Data, block.PrevHash)
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

// CreateBlock creates a new block in the blockchain.
func (bc *Blockchain) CreateBlock(data string) {
	prevBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
		Hash:      "",
	}
	newBlock.Hash = CalculateHash(newBlock)
	bc.Chain = append(bc.Chain, newBlock)
}

// NewBlockchain creates a new blockchain with a genesis block.
func NewBlockchain() *Blockchain {
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "Genesis Block",
		PrevHash:  "",
		Hash:      "",
	}
	genesisBlock.Hash = CalculateHash(genesisBlock)
	return &Blockchain{Chain: []Block{genesisBlock}}
}

func main() {
	// Create a new blockchain
	blockchain := NewBlockchain()

	// Add some blocks to the blockchain
	blockchain.CreateBlock("Block 1 Data")
	blockchain.CreateBlock("Block 2 Data")

	// Print the blockchain
	blockchainJSON, _ := json.MarshalIndent(blockchain, "", "  ")
	fmt.Println(string(blockchainJSON))
}
