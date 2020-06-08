package main

import "fmt"

func defer_test() {

	defer fmt.Println("World")
    fmt.Println("Hello")
}

func a() int {
    i := 0
    defer fmt.Println(i)
    i++
    return i
}

func a1() int {
    i := 0
    defer  func(i int) int {
	  return i
		}(i)
    i++
    return i
}

func test() {
	fmt.Println("counting")
	defer fmt.Println("counting 2")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}


func main() {
//defer simple
defer_test()
fmt.Println(a())
fmt.Println(a1())
//   fmt.Println("after:", k)
}