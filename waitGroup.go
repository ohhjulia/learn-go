package main

import (
	"fmt"
	"sync"
)

func main5() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(num int, function func()) {
			fmt.Println(num)
			function()
		}(i, wg.Done)
	}
	wg.Wait()
}
