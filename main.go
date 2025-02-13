package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/manimovassagh/go-clock/models"
)

type Blockchain struct {
	Chain      []models.Block
	Difficulty int
}

func createGenesisBlock() models.Block {
	return models.Block{
		Index: 0, Timestamp: time.Now().String(),
		Data:         "Genesis Block",
		PreviousHash: "", Hash: "", Nonce: 0}
}

func (bc *Blockchain) GetLatestBlock() models.Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *Blockchain) AddBlock(newBlock models.Block) {
	newBlock.PreviousHash = bc.GetLatestBlock().Hash
	log.Printf("Adding new block with index %d and previous hash %s", newBlock.Index, newBlock.PreviousHash)
	newBlock.MineBlockConcurrent(bc.Difficulty) // Use stable MineBlock function
	bc.Chain = append(bc.Chain, newBlock)
	log.Printf("Block added with hash %s", newBlock.Hash)
	bc.PrintBlockchain()
}

func (bc *Blockchain) IsChainValid() bool {
	for i := 1; i < len(bc.Chain); i++ {
		currentBlock := bc.Chain[i]
		previousBlock := bc.Chain[i-1]

		if currentBlock.Hash != currentBlock.CalculateHash() {
			return false
		}

		if currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

func (bc *Blockchain) PrintBlockchain() {
	blockchainJSON, err := json.MarshalIndent(bc, "", "  ")
	if err != nil {
		log.Println("Error marshalling blockchain:", err)
		return
	}
	fmt.Println(string(blockchainJSON))
	log.Println(string(blockchainJSON))
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Chain:      []models.Block{createGenesisBlock()},
		Difficulty: 7,
	}
}

func main() {
	// Set up logging to a file
	logFile, err := os.OpenFile("blockchain.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	block := models.Block{
		Index:        1,
		Timestamp:    time.Now().String(),
		Data:         "Block 1 Data",
		PreviousHash: "",
	}

	difficulty := 5 // Increase the difficulty level

	log.Println("Starting mining process...")
	block.MineBlockConcurrent(difficulty) // Use stable MineBlock function
	log.Println("Mining process completed")

	blockchain := NewBlockchain()
	blockchain.AddBlock(block)
	log.Println("Blockchain is valid:", blockchain.IsChainValid())
	blockchain.PrintBlockchain()
}
