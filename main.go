package main

import (
	"bytes"
	"crypto/sha256"
)

// Block is the main struct in this blockchain implementation
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// DeriveHash method to calculate block's hashes
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock does exact this...
func CreateBlock(data string, PrevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), PrevHash}
	block.DeriveHash()
	return block
}

func main() {

}
