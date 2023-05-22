package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	mut := sync.RWMutex{} //<== mut :=sync.Mutex{}

	var score []int
	wg.Add(4)
	go func() {
		fmt.Println("R one")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}()
	go func() {
		fmt.Println("R two")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}()
	go func() {
		fmt.Println("R three")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}()
	go func() {
		mut.RLock()
		fmt.Println("R read :", score) // Here score may or may not have value .It would be based on the execution .Its why its says as RWmutex{} is often would not get used .
		mut.RUnlock()
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Score is :", score)
}
