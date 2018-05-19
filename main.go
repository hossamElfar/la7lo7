package main

import (
	"fmt"
	"la7lo7/types"
	"strconv"
)

func main() {
	bc := types.Newtransactions()

	bc.AddBlock("Send 1 la7lo7 to 7oss")
	bc.AddBlock("Send 2 more la7lo7s to abo reda")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := types.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
