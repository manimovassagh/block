package models

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/fatih/color"
)

type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PreviousHash string
	Hash         string
	Nonce        int
}

func (b *Block) CalculateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PreviousHash + strconv.Itoa(b.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (b *Block) MineBlock(difficulty int) {
	startTime := time.Now() // Record the start time

	target := ""
	for i := 0; i < difficulty; i++ {
		target += "0"
	}
	b.Hash = b.CalculateHash() // Initialize the hash before entering the loop
	for b.Hash[:difficulty] != target {
		b.Nonce++
		b.Hash = b.CalculateHash()
	}
	endTime := time.Now() // Record the end time
	color.Green("************************************")
	color.Green("BLOCK MINED: %s", b.Hash)
	color.Red("Mining took: %v", endTime.Sub(startTime))
	color.Green("************************************")
}
