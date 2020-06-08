package main

import "fmt"

type myfun func(x, y int) int

func (f myfun) fun(x,y int) int {
	return f(x, y)
}

func f2(x, y int) int {
  return  x+ y
}

func f3() myfun {
  return func(x, y int) int{
    return x+ y
  }
}

func main(){
  fmt.Println(myfun(f2).fun(2,3)) //func type as receiver 
  fmt.Println(f3()(2,3)) //func type as return type
  fmt.Println(f3().fun(2,3))//func type as return type and call method on it
}