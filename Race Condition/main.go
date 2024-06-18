package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition :-")
	var score = []int{0}

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		mut.Lock()
		fmt.Println("1st")
		score = append(score, 1)
		wg.Done()
	}(wg, mut)

	wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("3rd")
		score = append(score, 3)
		wg.Done()
	}(wg, mut)

	wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("2nd")
		score = append(score, 2)
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
