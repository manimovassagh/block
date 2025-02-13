package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/manimovassagh/go-clock/models"
)

const DIFFICULTY = 7

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
	color.Cyan("Adding new block with index %d and previous hash %s", newBlock.Index, newBlock.PreviousHash)
	newBlock.MineBlock(bc.Difficulty) // Use stable MineBlock function
	bc.Chain = append(bc.Chain, newBlock)
	color.Green("Block added with hash %s", newBlock.Hash)
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
		color.Red("Error marshalling blockchain: %v", err)
		return
	}
	color.Yellow(string(blockchainJSON))
	log.Println(string(blockchainJSON))
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Chain:      []models.Block{createGenesisBlock()},
		Difficulty: DIFFICULTY,
	}
}

func main() {
	// Set up logging to a file
	logFile, err := os.OpenFile("blockchain.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		color.Red("Failed to open log file: %v", err)
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

	color.Blue("Starting mining process...")
	block.MineBlock(DIFFICULTY) // Use stable MineBlock function
	color.Blue("Mining process completed")

	blockchain := NewBlockchain()
	blockchain.AddBlock(block)
	color.Magenta("Blockchain is valid: %v", blockchain.IsChainValid())
	blockchain.PrintBlockchain()
}
