package main

import (
	"fmt"

	"github.com/lbaracat/golang-blockchain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("First block after The Beginning")
	chain.AddBlock("Second block after The Beginning - third in this blockchain")
	chain.AddBlock("Fourth block in this blockchain...")

	for k, block := range chain.Blocks {
		fmt.Printf("Block #%d\n", k)
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}
}
