package blockchain

//package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

// Block represents a single block in the blockchain.
type Block struct {
	Index     int    `json:"index"`
	Timestamp string `json:"timestamp"`
	Data      string `json:"data"`
	PrevHash  string `json:"prevhash"`
	Hash      string `json:"hash"`
}

// Blockchain is a simple blockchain structure.
type Blockchain struct {
	Chain []Block `json:"chain"`
}

// CalculateHash calculates the SHA-256 hash of a block.
func CalculateHash(block Block) string {
	data := fmt.Sprintf("%d%s%s%s", block.Index, block.Timestamp, block.Data, block.PrevHash)
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

// CreateBlock creates a new block in the blockchain.
func CreateBlock(data string, prevHash string) Block {
	newBlock := Block{
		Index:     0, // This will be set to the correct index when adding to the blockchain
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevHash, // This is the hash of the last block in the blockchain
		Hash:      "",       // This will be set after the block is created
	}

	// The hash of the new block should be calculated with all the block data including the PrevHash
	newBlock.Hash = CalculateHash(newBlock)

	// Now return the new block, ready to be added to the blockchain
	return newBlock
}

// NewBlockchain creates a new blockchain with a genesis block.
func NewBlockchain() Blockchain {
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "Genesis Block",
		PrevHash:  "",
		Hash:      "",
	}
	genesisBlock.Hash = CalculateHash(genesisBlock)
	chain := Blockchain{
		Chain: []Block{genesisBlock},
	}

	return chain

}

func main() {
	// Create a new blockchain
	blockchain := NewBlockchain()

	// Add some blocks to the blockchain
	//blockchain.CreateBlock("Block 1 Data")
	//blockchain.CreateBlock("Block 2 Data")
	// Print the blockchain
	blockchainJSON, _ := json.MarshalIndent(blockchain, "", "  ")
	fmt.Println(string(blockchainJSON))
}
