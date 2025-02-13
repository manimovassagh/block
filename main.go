package main

import (
	"fmt"
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
	newBlock.MineBlock(bc.Difficulty) // Update function call
	bc.Chain = append(bc.Chain, newBlock)
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

func NewBlockchain() *Blockchain {
	return &Blockchain{
		Chain:      []models.Block{createGenesisBlock()},
		Difficulty: 4,
	}
}

func main() {
	block := models.Block{
		Index:        1,
		Timestamp:    time.Now().String(),
		Data:         "Block 1 Data",
		PreviousHash: "",
	}

	difficulty := 5 // Increase the difficulty level

	fmt.Println("Mining block with concurrency...")
	block.MineBlock(difficulty) // Update function call

	block.Nonce = 0 // Reset nonce for fair comparison
	fmt.Println("Mining block without concurrency...")
	block.MineBlockOld(difficulty)
}
