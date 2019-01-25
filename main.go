package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func goroutines(i int) {
	time.Sleep(time.Second)
	fmt.Printf("current i = %d \n", i)
	wg.Done()
}

func main() {
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go goroutines(i)
	}
	wg.Wait()
}
