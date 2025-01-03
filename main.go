package main

import (
	"fmt"
	"encoding/binary"
	"crypto/sha256"
	"time"
)


type block struct {

	prevBlockHash 	[32]byte
	blockIndex 		uint64
	timestamp 		string
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
	allBytes = append(allBytes, []byte(myBlock.timestamp)...)
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


func verifyBlock(myBlock block) bool {

	if hashBlock(myBlock)[0] == 0 && hashBlock(myBlock)[1] == 0 {
		return true
	}
	
return false

}

func findNonce(myBlock block) uint64 {

	myBlock.nonce = 0

	for hashBlock(myBlock)[0] != 0 || hashBlock(myBlock)[1] != 0 {
		myBlock.nonce += 1
	}

	return myBlock.nonce
	
}


func main(){

// Declare chain

	var chain blockChain

// Init and seed Genisis Block
	genisisBlock := block{[32]byte{}, 0, time.Time.String(time.Now()), 0 }	
	genisisBlock.nonce = findNonce(genisisBlock)
	
	if verifyBlock(genisisBlock){
		chain.blocks = append(chain.blocks, genisisBlock)
		fmt.Printf("The Genisis Block:     %x, Block index: %d, Time: %s, Nonce: %d \n", genisisBlock.prevBlockHash, genisisBlock.blockIndex, genisisBlock.timestamp, genisisBlock.nonce)
	}
	

	for i := 1; i <= 100; i++ {
	
		newblock := block{hashBlock(chain.blocks[i-1]), uint64(i), time.Time.String(time.Now()), 0}	
		newblock.nonce = findNonce(newblock)

		if verifyBlock(newblock){
			chain.blocks = append(chain.blocks, newblock)	
			fmt.Printf("Previous Block's hash: %x, Block index: %d, Time: %s, Nonce: %d\n", chain.blocks[i].prevBlockHash, chain.blocks[i].blockIndex,  chain.blocks[i].timestamp, chain.blocks[i].nonce)

		}
	}

	// verify the chain

	if verifyChain(chain){
		fmt.Println("\nBlockchain is validated")
	} else {
		fmt.Println("Blockchain is invalid")
	}
	
}
