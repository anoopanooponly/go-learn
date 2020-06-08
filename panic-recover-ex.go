package main

import ("fmt"
 "runtime/debug"
)
func recoverName(){

if r:= recover(); r != nil {
	fmt.Println("Recovered from :", r)
	debug.PrintStack()
}

}

func fullname(firstname *string, lastname *string) {

defer recoverName()
if firstname == nil {
	panic("runtime error: first name cannot be nil")
}
if lastname == nil {
	panic("runtime error: last name cannot be nil")
}
fmt.Printf("%s %s\n", *firstname, *lastname)
fmt.Println("returned normally from fullName")

}


func main() {

	s:= "anoop"




	fullname(&s, nil)
	fmt.Println("end of main")
}