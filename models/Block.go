package models

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"sync"
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

func (b *Block) MineBlock(difficulty int) {
	startTime := time.Now() // Record the start time

	target := ""
	for i := 0; i < difficulty; i++ {
		target += "0"
	}

	var wg sync.WaitGroup
	found := make(chan bool)
	var mu sync.Mutex
	var once sync.Once

	for i := 0; i < 4; i++ { // Create 4 goroutines for concurrent mining
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-found:
					return
				default:
					mu.Lock()
					b.Nonce++
					b.Hash = b.CalculateHash()
					mu.Unlock()
					if b.Hash[:difficulty] == target {
						endTime := time.Now() // Record the end time
						fmt.Println("************************************")
						fmt.Println("BLOCK MINED: ", b.Hash)
						fmt.Println("Mining took: ", endTime.Sub(startTime))
						fmt.Println("************************************")
						once.Do(func() { close(found) })
						return
					}
				}
			}
		}()
	}

	wg.Wait()
}
