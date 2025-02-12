package main

import (
	"encoding/json"
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
	newBlock.MineBlock(bc.Difficulty)
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
	bc := NewBlockchain()

	fmt.Println("Mining block 1...")
	bc.AddBlock(models.Block{Index: 1,
		Timestamp: time.Now().String(),
		Data:      "Block 1 Data", PreviousHash: "",
		Hash: "", Nonce: 0})

	fmt.Println("Mining block 2...")
	bc.AddBlock(models.Block{Index: 2,
		Timestamp: time.Now().String(),
		Data:      "Block 2 Data", PreviousHash: "",
		Hash: "", Nonce: 0})

	fmt.Println("Blockchain valid?", bc.IsChainValid())

	blockchainJSON, _ := json.MarshalIndent(bc, "", "  ")
	fmt.Println(string(blockchainJSON))
}
