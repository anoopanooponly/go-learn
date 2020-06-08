package main


import "fmt"


var a [5]int

func main(){
	a[2] = 3

	b := [5]int{1,2,3,4, 5}
	fmt.Print(a, len(a))
	fmt.Println(b, len(b))

	

	a1 := [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
	fmt.Println(a1)


	//Slice not array. Array always have predefined size
	slice_not_arr := []int {1,2,3}
	fmt.Println(slice_not_arr)
}