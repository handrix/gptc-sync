package syncer

import "fmt"

func GetBlock() {
	currentBlock := CurrentBlock()
	fmt.Println(currentBlock)
}
