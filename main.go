package main

import (
	"fmt"
	"encoding/binary"
	"crypto/sha256"
	"math/rand"
	"time"
)


type block struct {

	prevBlockHash 	[32]byte
	blockIndex 		uint64
	nonce			uint64
	
}

type blockChain struct {
	blocks []block
}

type transaction struct {
	sender 			string
	recipient 		string
	ammount			int
}

func uint64ToByteSlice(x uint64) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, x)
	return bytes
}

func uint32ToByteSlice(x uint32) []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, x)
	return bytes
}

func blockToByteSlice(myBlock block) []byte {

	allBytes := []byte{}
	allBytes = append(allBytes, myBlock.prevBlockHash[:]...)
	allBytes = append(allBytes, uint64ToByteSlice(myBlock.blockIndex)...)
	allBytes = append(allBytes, uint64ToByteSlice(myBlock.nonce)...)

	return allBytes	
	
}

func hashBlock(myBlock block) [32]byte {
	sum := sha256.Sum256(blockToByteSlice(myBlock))
	return sum
}

func verifyChain(chain blockChain) bool {

	for i := 0; i < len(chain.blocks)-1; i++ {
		if hashBlock(chain.blocks[i]) != chain.blocks[i+1].prevBlockHash{
			return false
		}
	}
	
	return true
}



func main(){

// Declare chain

	var chain blockChain

// Init and seed Genisis Block
	x := [32]byte{}	
	rand.Seed(time.Now().UnixNano())
	
	for i := 0; i < 32; i++ {
		x[i] = byte(rand.Intn(255))
	}
	
	x[27] = 'r'
	x[28] = 'h'
	x[29] = 'e'
	x[30] = 't'
	x[31] = 't'
	
	genisisBlock := block{x, 0, 0 }	
	chain.blocks = append(chain.blocks, genisisBlock)


	for i := 0; i <= 1000; i++ {
		newblock := block{hashBlock(chain.blocks[i]), uint64(i+1), 0}	
		chain.blocks = append(chain.blocks, newblock)	
		fmt.Printf("Previous Block hash: %x, Block index: %d, Nonce: %d\n", chain.blocks[i].prevBlockHash,chain.blocks[i].blockIndex,chain.blocks[i].nonce)
	}

	if verifyChain(chain){
		fmt.Println("\nBlockchain is validated")
	} else {
		fmt.Println("Blockchain is invalid")
	}
	
}
