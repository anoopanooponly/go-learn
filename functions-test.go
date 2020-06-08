package main

import "fmt"

//ex for func, variadic func, closure, recursion


//closure and anonymous func

func intIter() func() int {
	i := 0
	// The returned function closes over the variable i to form a closure.
	return func() int {
		i++ 
		return i 
	}
}

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func print(arr ...int){
	for _, v:= range arr {
		fmt.Println(v)
	}
}

func main() {
	print(1)
	print(3,4,5)
	arr := []int {7,8,9,10}
	print(arr...)

	//closure
	intIterNext := intIter()
	fmt.Println(intIterNext())
	fmt.Println(intIterNext())
	fmt.Println(intIterNext())

	fmt.Println(fact(7))

}