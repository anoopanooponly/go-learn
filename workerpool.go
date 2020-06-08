package main

import ("fmt"
"time"
"sync"
"math/rand"
)

type Job struct {
	id int
	randomnum int
}

type Result struct {
	job Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results =make(chan Result, 10)

//job to find sum of given digits
func processJob(digits int) int {

	sum := 0
	var x = 0
	for digits != 0 {
	
		x = digits%10
		sum += x
		digits = digits/10
    
	}
	time.Sleep(2 * time.Second)
	return sum
}

func createJobs(noofjobs int) {

	for i:=0;i < noofjobs;i++ {
		randomno := rand.Intn(999)
        job := Job{i, randomno}
        jobs <- job
	}

	close(jobs)
}

func worker(wg *sync.WaitGroup) {

	
	for job := range jobs {
		output := Result{job, processJob(job.randomnum)}
		results <- output
	}

	wg.Done()
}

func createWorkerPool(noofworkers int) {
	var wg sync.WaitGroup
	for i:=0;i<noofworkers;i++ {
	  wg.Add(1)
	  go worker(&wg)
	  
	}
	
	wg.Wait()
	close(results)
	
	
	
}

func result(done chan bool) {  
    for result := range results {
        fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomnum, result.sumofdigits)
    }
    done <- true
}

func main(){

	//create 100's of jobs and n  number of workers to work on it
	//create jobs and assign jobs to jobs channel. so create two channels for jobs and results globally
	
	startTime := time.Now()
    noOfJobs := 100
	go createJobs(noOfJobs)
	

	 done := make(chan bool)
	 go result(done) //results reading need to be in gorutine because statements in main will be 
	
	noOfWorkers := 50
	createWorkerPool(noOfWorkers)
	
	
	 <- done
	// for i:=0; i < noOfJobs;i++ {
	// 	x:= rand.Intn(999)
		
	// 	fmt.Printf("input random no %d , sum of digits %d\n", x, processJob(x))
	// }
	endTime := time.Now()
    diff := endTime.Sub(startTime)
    fmt.Println("total time taken ", diff.Seconds(), "seconds")
}