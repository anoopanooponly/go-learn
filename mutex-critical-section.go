package main

import ("fmt"
"sync"
)

var x = 0
func incrementNormal(wg *sync.WaitGroup){
	x = x + 1
	wg.Done()
}

func incrementWithMutex(wg *sync.WaitGroup, m *sync.Mutex){
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

// we have created a buffered channel of capacity 1 and this is passed to the increment 
// Goroutine in line no. 18. This buffered channel is used to ensure that only one Goroutine
//  access the critical section of code which increments x. 
func channelForBlocking(wg *sync.WaitGroup, c chan bool) {
	c <- true
	x = x + 1
	<- c
	wg.Done()
}

func main(){

	var wg sync.WaitGroup
	// var m sync.Mutex
    c := make(chan bool, 1)
	for i:=0;i< 1000;i++ {
		wg.Add(1)
		// go incrementNormal(&wg)
		// go incrementWithMutex(&wg, &m)
		go channelForBlocking(&wg,c)
	}

	wg.Wait()
	fmt.Println(x)
}