package syncer

import (
	"fmt"
	"gptc-sync/models"
	"time"
)

var blockMethod = BlockMethod{}

func publisher(task chan int64) {
	if len(task) < 1 {
		maxIndex := models.MaxIndex()
		targetIndex := maxIndex + 100
		for i := maxIndex; i <= targetIndex; i++ {
			task <- i
		}
	}
}

func worker(number int64) {
	result := blockMethod.BlockDetail(number)
	v := models.GetBlock(number)
	// 不存在则创建
	if len(v) == 0 {
		models.AddBlock(result)
		fmt.Println(result["numberInt"])
		// 已存在则跳过
	} else {
		return
	}
}

func SyncBlock() {
	// 定时器，定时触发publisher，避免忙循环
	tick := time.Tick(1e8)
	task := make(chan int64)

	for {
		select {
		case <-tick:
			go publisher(task)
		case i := <-task:
			go worker(i)
		}
	}
}
