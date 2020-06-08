package main


import "fmt"


func passbyval(i int) {
	i++
}

func passbyref(i *int) {
	*i++
}

func main(){
	var a int = 0
	fmt.Println("before:",a)
	passbyval(a)
	fmt.Println("After:",a)
	passbyref(&a)
	fmt.Println("after ref:",a)

	fmt.Println("pointer:", &a)
}