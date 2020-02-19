package main

import (
	"fmt"
	"sync"
	"time"
)

//Five silent philosophers sit at a round table with bowls of spaghetti. Forks are placed between each pair of adjacent
//philosophers. Each philosopher must alternately think and eat. However, a philosopher can only eat spaghetti when they
//have both left and right forks. Each fork can be held by only one philosopher and so a philosopher can use the fork only
//if it is not being used by another philosopher. After an individual philosopher finishes eating, they need to put down
//both forks so that the forks become available to others. A philosopher can take the fork on their right or the one on
//their left as they become available, but cannot start eating before getting both forks. Eating is not limited by the
//remaining amounts of spaghetti or stomach space; an infinite supply and an infinite demand are assumed. The problem is
//how to design a discipline of behavior (a concurrent algorithm) such that no philosopher will starve; i.e., each can
//forever continue to alternate between eating and thinking, assuming that no philosopher can know when others may want to
//eat or think.

func dine(p string, lFork, rFork *sync.Mutex, wg *sync.WaitGroup) {
	sTime := time.Second/100
	fmt.Println(p, "seated")
	fmt.Println(p, "hungry")
	lFork.Lock()
	rFork.Lock()
	fmt.Println(p, "eating")
	time.Sleep(sTime)
	lFork.Unlock()
	rFork.Unlock()
	fmt.Println(p, "thinking")
	time.Sleep(sTime)
	fmt.Println(p, "satisfied")
	wg.Done()
	fmt.Println(p, "left table")

}
func main(){
	ph := []string{"a", "b", "c", "d", "e"}
	fmt.Println("empty table")
	var wg sync.WaitGroup
	wg.Add(len(ph))
	fork := &sync.Mutex{}
	lFork := fork

	for i := 0; i< len(ph); i++ {
		rFork := &sync.Mutex{}
		go dine(ph[i], lFork, rFork, &wg)
		lFork = rFork
	}

	wg.Wait()
	fmt.Println("empty table")
}