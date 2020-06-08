package main

import ( "fmt"
"time"
)


func server1(c chan string){

	time.Sleep(2 * time.Second)
	c <- "from server1"
}

func server2(c chan string) {
	time.Sleep(4 * time.Second)
	c <- "from server2"
}

func main() {

	s1:= make(chan string)
	s2:= make(chan string)

	go server1(s1)
	go server2(s2)

	for {
		select {

		case x:= <- s1:
			fmt.Println("server1:", x)
			return
		case y:= <- s2:
			fmt.Println("server2:", y)
			return
		default:	
		fmt.Println("no server")
		}
	}
	
}