package main

import ("fmt"
"time")

func main(){

	//normal switch
	for i:=0; i < 3; {
		i++
		switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		default:
			fmt.Println(time.Now().Weekday())
		}
	}
	
	//

	switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
	}
	//switch without cond
 t:= time.Now()
	switch {

	case t.Hour() < 12:
		fmt.Println("Forenoon")
	default:
		fmt.Println("afternoon")
	}

	whatamI := typeof
fmt.Print(whatamI("abc"))	

	
}

func typeof(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	
	default:
		return "unknown"
	}
}