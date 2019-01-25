package main

import (
	"fmt"
	"math/rand"
	"time"
)

func runFuncFile() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Hello World!!", rand.Intn(100))
}
