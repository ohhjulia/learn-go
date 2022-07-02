package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main1() {

	count := uint32(0)

	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}
	// sign := make(chan struct{}, 10)

	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
			// fmt.Println(i)
			// sign <- struct{}{}
		}(i)
	}

	trigger(10, func() {})

	// for i := 0; i < 10; i++ {
	// 	<-sign
	// }
	// time.Sleep(10)

}
