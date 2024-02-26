package blockchain

//package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
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
func CreateBlock(data string) Block {
	//Gets previous block data
	var chain Blockchain
	jsonChainData, err := os.ReadFile("chain.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(jsonChainData, &chain)
	if err != nil {
		return Block{}
	}

	newBlock := Block{
		Index:     chain.Chain[len(chain.Chain)-1].Index + 1,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  chain.Chain[len(chain.Chain)-1].Hash,
		Hash:      "",
	}
	newBlock.Hash = CalculateHash(newBlock)
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
