package blockchain

import (
	"bytes"
	"testing"
)

func Test(t *testing.T) {
	var blockData string = "test block 1"

	var exceptedBlockChainData = []struct {
		data     []byte
		hash     []byte
		prevHash []byte
		nonce    int
	}{
		{
			[]byte{73, 110, 32, 116, 104, 101, 32, 98, 101, 103, 105, 110, 110, 105, 110, 103, 32, 100, 101, 118, 32, 99, 114, 101, 97, 116, 101, 100, 32, 116, 104, 105, 115, 32, 98, 108, 111, 99, 107, 32, 97, 110, 100, 32, 116, 104, 101, 32, 98, 108, 111, 99, 107, 99, 104, 97, 105, 110},
			[]byte{0, 11, 249, 149, 56, 103, 157, 29, 150, 83, 196, 5, 155, 172, 191, 67, 193, 230, 111, 10, 146, 250, 90, 196, 119, 76, 97, 155, 91, 237, 253, 16},
			[]byte{},
			1080,
		},
		{
			[]byte{116, 101, 115, 116, 32, 98, 108, 111, 99, 107, 32, 49},
			[]byte{0, 2, 190, 2, 243, 134, 65, 173, 41, 255, 200, 173, 188, 50, 43, 90, 222, 207, 25, 233, 24, 233, 236, 109, 3, 79, 130, 237, 38, 119, 84, 77},
			[]byte{0, 11, 249, 149, 56, 103, 157, 29, 150, 83, 196, 5, 155, 172, 191, 67, 193, 230, 111, 10, 146, 250, 90, 196, 119, 76, 97, 155, 91, 237, 253, 16},
			750,
		},
	}

	var toHex = []struct {
		myInt64 int64
		want    []byte
	}{
		{4611756388245323776, []byte{64, 0, 64, 0, 64, 0, 64, 0}},
	}

	for _, v := range toHex {
		got := ToHex(v.myInt64)
		if bytes.Compare(got, v.want) != 0 {
			t.Errorf("ToHex(%v) == %v, want %v", v.myInt64, got, v.want)
		}
	}

	testChain := InitBlockChain()
	testChain.AddBlock(blockData)

	for k, block := range testChain.Blocks {
		if bytes.Compare(block.Data, exceptedBlockChainData[k].data) != 0 {
			t.Errorf("Data error in block %v: Have (%v), wants (%v)", k, block.Data, exceptedBlockChainData[k].data)
		}
		if bytes.Compare(block.Hash, exceptedBlockChainData[k].hash) != 0 {
			t.Errorf("Hash error in block %v: Have (%v), wants (%v)", k, block.Hash, exceptedBlockChainData[k].hash)
		}
		if bytes.Compare(block.PrevHash, exceptedBlockChainData[k].prevHash) != 0 {
			t.Errorf("Pevious hash error in block %v: Have (%v), wants (%v)", k, block.PrevHash, exceptedBlockChainData[k].prevHash)
		}
		if block.Nonce != exceptedBlockChainData[k].nonce {
			t.Errorf("NOnce error in block %v: Have (%v), wants (%v)", k, block.Nonce, exceptedBlockChainData[k].nonce)
		}

		pow := NewProof(block)

		if !pow.Validate() {
			t.Errorf("pow.Validate() error in block %v", k)
		}
	}
}
