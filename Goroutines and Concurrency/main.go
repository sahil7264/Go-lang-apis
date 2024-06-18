// Concurrency -> Multiple tasks but not at same time
// Parallelism -> Multiple tasks at same time

// Go routines -> This is  How Parallelism implemented
// Thread -> Managed by OS,  Fixed Task
// Go routines -> Managed by Go runtime , Flexible Stack

// Go routines-> NOt Communicate  by sharing  Memory bu Share memory by communicating

// Methods to implement
// 1) time.Sleep()
// 2) sync.waitGroup() package -> problem when multiple goroutine tries to write on same memory
// 3) Mutex locks
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func greet(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

var signals = []string{"test"}
var wg sync.WaitGroup // pointer
var mut sync.Mutex    // pointer

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("IN ENDPOINT")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d Status Code for %s \n", res.StatusCode, endpoint)
	}
}

func main() {

	//Method 1
	// go greet("sahil") // go keyword -> NO waiting then not print
	// greet("shile")

	//Method 2
	websiteList := []string{
		"https://google.com",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://leetcode.com",
	}
	for _, web := range websiteList {
		//using package sync
		go getStatusCode(web)
		wg.Add(1) // adding number that are adding
	}
	wg.Wait() // not allowing to finish
	fmt.Println(signals)
	//Method 3
}
