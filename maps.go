package main

import "fmt"

func testMap() {
	// var mymap map[string]int
	mymap := make(map[string]int)
	mymap["hello"] = 45
	mymap["world"] = 67
	fmt.Println(mymap)
	fmt.Println("Hello world")
}
