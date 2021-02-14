package main

import (
	"fmt"
	"strconv"

	"github.com/andre-karrlein/blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First test")
	chain.AddBlock("Second test")
	chain.AddBlock("Third test")

	for _, block := range chain.Blocks {
		//fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
