package main

import (
	"fmt"
	"errors")

func main() {
  
	x := []int {1,2,3,4, 5, 6}
	if a, err := f1(x...); err != nil {
     fmt.Println("Error", a)
	} else {
		fmt.Println("no issue ", a)
	}

	x1 := []int {1,2,3,4, 5, 6}
	if a, err := f2(x1...); err != nil {
     fmt.Println("Error", err)
	} else {
		fmt.Println("no issue ", a)
	}


	_, e := f2(5)
	
    if ae, ok := e.(MyCustError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}


func f1(x ...int) (int , error) {
  
	for _,b := range x {
		if b == 5 {
			return b, errors.New("can't work with 5")
		}
	}
	return -1, nil
}

type MyCustError struct {
	arg int
	prob string
}

func (e MyCustError) Error() string {
	return fmt.Sprintf("%d -> %s", e.arg, e.prob)
}


func f2(x ...int) (int , error) {
  
	for _,b := range x {
		if b == 5 {
			return b, MyCustError{b, "Custom Error"} // impements error interface and return, so
			//    MyCustError should have Error() method implemented 
		}
	}
	return -1, nil
}
