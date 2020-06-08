package main

import "fmt"


func main() {
	m := make(map[string] int)
	m["k"] = 2
	m["l"] = 3
	fmt.Println(m)

	m1 := make(map[byte] int)
	m1['a'] = 2
	fmt.Println(m1['a'])

	fmt.Println("len:", len(m))

	delete(m, "l")
	fmt.Println(m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs) // To check if key present in map
	
	_, prs = m["k"]
	fmt.Println("prs:", prs)
	
	mapinit := map[string] int { "a": 1, "b": 2, "c": 3}
	
	fmt.Print(mapinit) 

	//range example

	for k,v := range mapinit {
         fmt.Printf("%s -> %v \n", k,v)
	}

}