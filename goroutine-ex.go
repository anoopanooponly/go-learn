package main

import (
	
	"fmt"
	"time"
)

	func test(s string) {

		for i:= 0;i < 3;i++ {
			fmt.Printf("%v -> %v\n", i, s)
		}
	}
	func main() {
	fmt.Println("starrt")
	   go test("abv")

	   go func(msg string) {
        fmt.Println(msg)
    }("going")
	   time.Sleep(time.Second)
	   fmt.Println("end")
	} 