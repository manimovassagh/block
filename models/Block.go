package models

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
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
type miningResult struct {
	nonce int
	hash  string
}

// Stable MineBlock function without concurrency
// func (b *Block) MineBlock(difficulty int) {
// 	startTime := time.Now() // Record the start time

// 	target := ""
// 	for i := 0; i < difficulty; i++ {
// 		 target += "0"
// 	}
// 	b.Hash = b.CalculateHash() // Initialize the hash before entering the loop
// 	for b.Hash[:difficulty] != target {
// 		b.Nonce++
// 		fmt.Println("Nonce: ", b.Nonce)
// 		b.Hash = b.CalculateHash()
// 	}
// 	endTime := time.Now() // Record the end time
// 	fmt.Println("************************************")
// 	fmt.Println("BLOCK MINED: ", b.Hash)
// 	fmt.Println("Mining took: ", endTime.Sub(startTime))
// 	fmt.Println("************************************")
// }

func (b *Block) MineBlockConcurrent(difficulty int) {
	startTime := time.Now()
	target := strings.Repeat("0", difficulty)

	// Check if current hash already meets target
	if initialHash := b.CalculateHash(); initialHash[:difficulty] == target {
		b.Hash = initialHash
		fmt.Println("Block already valid")
		return
	}

	numWorkers := runtime.NumCPU()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resultChan := make(chan miningResult, 1)
	var foundFlag atomic.Bool

	originalNonce := b.Nonce

	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			nonce := originalNonce + 1 + workerID
			for {
				select {
				case <-ctx.Done():
					return
				default:
					if foundFlag.Load() {
						return
					}

					// Create temporary block with current nonce
					tempBlock := *b
					tempBlock.Nonce = nonce
					hash := tempBlock.CalculateHash()

					if hash[:difficulty] == target {
						if foundFlag.CompareAndSwap(false, true) {
							resultChan <- miningResult{nonce: nonce, hash: hash}
							cancel()
						}
						return
					}

					nonce += numWorkers
					fmt.Println("Worker", workerID, "tried nonce", nonce)
				}
			}
		}(i)
	}

	// Wait for result
	result := <-resultChan
	b.Nonce = result.nonce
	b.Hash = result.hash
	fmt.Println(result)

	endTime := time.Now()
	fmt.Println("************************************")
	fmt.Println("BLOCK MINED:", b.Hash)
	fmt.Println("Mining took:", endTime.Sub(startTime))
	fmt.Println("************************************")
}