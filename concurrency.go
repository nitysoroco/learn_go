package main

import (
	"fmt"
	"sync"
	"time"
)

// Concurrency: Dealing with multiple things at a time
// Parrellelism: Doing multiple things at a time
// WaitGroup, Add(1), Wait(), Done()

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

func runConcurrency() {

	wg.Add(1)
	go say("Hello")
	wg.Add(1)
	go say("World")
	wg.Add(1)
	go say("Turbo")
	wg.Wait()
	// time.Sleep(time.Second * 10)
}
