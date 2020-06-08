package main

import ("fmt"
"time"
)

func producer(datas chan int) {

	for i:=0; i< 20;i++ {
		datas <- i
		fmt.Println("producing :", i)
	}

	close(datas)
}

func consumer(datas chan int, done chan bool){

	for data:= range datas {
		fmt.Println("Consuming :", data)
	} 
	done <- true
}

func main(){

//   var datas = make (chan int, 10)
//   var done = make (chan bool)
//   go producer(datas)
  
//   go consumer(datas, done)
//   <- done

//   fmt.Println("Done")

// s := []int{7, 2, 8, -9, 4, 0}

// c := make(chan int)

// go sum(s[:len(s)/2],c)
// go sum(s[len(s)/2:],c)

// x,y := <-c, <-c

// fmt.Println(x+y)

ch := make(chan int, 2)
    go write(ch)
    //
    for v := range ch {
        fmt.Println("read value", v,"from ch")
        time.Sleep(2 * time.Second)

    }

}

func write(ch chan int) {  
    for i := 0; i < 15; i++ {
		ch <- i
		time.Sleep(2 * time.Second)
        fmt.Println("successfully wrote", i, "to ch")
    }
    close(ch)
}

func sum(arr []int, c chan int) {

	sum := 0
	for _, v := range arr {
		sum += v
	}
	c <- sum // send sum to c
}