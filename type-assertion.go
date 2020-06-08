package main

import "fmt"

//  t:= t.(T) A type assertion provides access to an interface values underlying concrete value
func main() {

	var s interface {} = "Hello"
	
	s1 := s.(string)
	fmt.Println(s1)

	// s2  := s.(int) //panic as s is not of type int, if used with 
	// second boolean argument panic wont be thrown
	// fmt.Println(s2)

	i, ok  := s.(int) //panic as s is nt of type int, if used with 
	// second boolean argument panic wont be thrown
	fmt.Println(i, ok)

	fmt.Println(typeof(32))

	//type assertion pointer
	var s3 = 20
	var addrs = &s3
	fmt.Println(s3)
	fmt.Println(*addrs)
	interType(addrs)
	
	
}

func typeof(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case *float64:
		return "float64 pointer"
	case *string:
		return "string pointer"
	default:
		return "unknown"
	}
}

func interType(i interface{}){
	t := i.(*int) // type assertion pointer
	fmt.Println(t)
}