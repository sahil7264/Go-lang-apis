package main

import (
	"fmt"
	"sync"
)

// Handoff to go for handling better
// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those values into another goroutine.

func main() {
	fmt.Print("Channels in GO : ")
	myCh := make(chan int, 2)

	// myCh <- 5 // arrow pointing toward left
	// fmt.Println(<-myCh)

	// Channel allow to send value when somebody is listening to him
	// No Listening -> No send value

	//Solution is GO routines

	wg := &sync.WaitGroup{}
	wg.Add(2)

	// ->  Receive  only goroutine
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// fmt.Println(<-myCh) // one listener can listen one valie
		// fmt.Println(<-myCh) // use to listen 2nd value , but to avoid we can use buffer

		// To avoid zero error
		val, isChannelOpen := <-myCh
		fmt.Println(val, isChannelOpen) // false -> only by close else true
		wg.Done()
	}(myCh, wg)

	// <-  send only goroutine
	go func(ch chan<- int, wg *sync.WaitGroup) {
		// close(myCh) // panic mode channel is close still you are sending
		// myCh <- 5 // one sender
		// myCh <- 6 // one sender send to one receiver , by default , have to give number when make is used
		// close(myCh)

		// listening only and channel closed then 0 get send as signal
		// while sending 0, cant say it is message or signal value

		myCh <- 0   // true
		close(myCh) // without send false;
		wg.Done()
	}(myCh, wg)
	wg.Wait()

}
