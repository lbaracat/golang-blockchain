package blockchain

// Block is the blockchain unit in this implementation
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// BlockChain struct for this implementation
type BlockChain struct {
	Blocks []*Block
}

// CreateBlock does exact this...
func CreateBlock(data string, PrevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), PrevHash, 0}
	pow := NewProof(block)

	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// AddBlock call CreateBlock and append this on BlockChain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)

}

// Genesis generates the first block
func Genesis() *Block {
	return CreateBlock("In the beginning dev created this block and the blockchain", []byte{})
}

// InitBlockChain makes the blockchain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
