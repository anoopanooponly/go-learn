package main

import "fmt"


func main(){

	//slice literals
	letters := []string{"a", "b", "c", "d"}
	fmt.Println(letters)

	x := letters[0:2]
	x[1] = "3" // a new slice referencing the old slice.
	// Slicing does not copy the slice's data. 
	// It creates a new slice value that points to the original array.
	fmt.Println("letters:", letters[1])

	s := make([]string, 5)
	s[0] = "a"
    s[1] = "b"
    s[2] = "c"
	s[3] = "d"
	s[4] = "f"
	fmt.Println(s)
	fmt.Println("len:", len(s))
	fmt.Println("Cap:", cap(s))
	s = append(s, "g","h")
	// s = append(s, "h")
	fmt.Println("len:", len(s))
	fmt.Println("Cap:", cap(s))

	s1 := s[0: 4]
	fmt.Println("s1:", s1)
    fmt.Println("len s1:", len(s1))
	fmt.Println("Cap s1:", cap(s1))

	fmt.Printf("0: %v,1: %v,2: %v,3: %v,4: %v,5: %v",s[0],s[1],s[2],s[3],s[4], s[5])

	c := make([]string, len(s))
    copy(c, s)
	fmt.Println("cpy:", c)
	
	//range

	for i,v := range c {
		fmt.Printf("%v -> %v",i,v)
	} 
}