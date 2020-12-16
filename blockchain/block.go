package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// Block is the blockchain unit in this implementation
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// BlockChain struct for this implementation
type BlockChain struct {
	blocks []*Block
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

// AddBlock call CreateBlock and append this on BlockChain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)

}

// Genesis generates the first block
func Genesis() *Block {
	return CreateBlock("In the beginning dev created this block and the blockchain", []byte{})
}

// InitBlockChain makes the blockchain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}