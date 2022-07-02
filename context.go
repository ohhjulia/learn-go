package main

import (
	"context"
	"fmt"
	"os"
	"sync/atomic"
)

func main6() {
	total := 12
	var num int32
	fmt.Printf("The number: %d [with context.Context]\n", num)
	cxt, cancelFunc := context.WithCancel(context.Background())
	for i := 1; i <= total; i++ {
		go addNum(&num, i, func() {
			if atomic.LoadInt32(&num) == int32(total) {
				cancelFunc()
			}
		})
	}
	<-cxt.Done()
	fmt.Println("End.")
	os.Stdout.Sync()
}

// addNum 用于原子地增加一次numP所指的变量的值。
func addNum(numP *int32, id int, deferFunc func()) {
	defer deferFunc()
	// for i := 0; ; i++ {
	// currNum := atomic.LoadInt32(numP)
	atomic.AddInt32(numP, 1)
	fmt.Println(*numP)

	// newNum := currNum + 1
	// time.Sleep(time.Millisecond * 200)
	// if atomic.CompareAndSwapInt32(numP, currNum, newNum) {
	// 	fmt.Printf("The number: %d [%d-%d]\n", newNum, id, i)
	// 	break
	// } else {
	// 	//fmt.Printf("The CAS operation failed. [%d-%d]\n", id, i)
	// }
	// }
}
