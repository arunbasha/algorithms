// Assume a cigarette requires three ingredients to make and smoke: tobacco, paper, and matches. There are three smokers
// around a table, each of whom has an infinite supply of one of the three ingredients — one smoker has an infinite
// supply of tobacco, another has paper, and the third has matches. A fourth party, with an unlimited supply of
// everything, chooses at random a smoker, and put on the table the supplies needed for a cigarrette. The chosen smoker
// smokes, and the process should repeat indefinitely.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	paper = iota
	grass
	match
)

var smokeMap = []string{"paper", "tobacco", "match"}

var smokers = []string{"Sandy","Apple","Daisy"}

type Table struct {
	paper   chan int
	tobacco chan int
	match   chan int
}

func arbitrate(t *Table, smokers [3]chan int, wg *sync.WaitGroup) {
	for {
		time.Sleep(time.Millisecond * 500)
		next := rand.Intn(3)
		fmt.Printf("Table chooses %s: %p\n", smokeMap[next], smokers[next])
		switch next {
		case paper:
			t.tobacco <- 1
			t.match <- 1
		case grass:
			t.paper <- 1
			t.match <- 1
		case match:
			t.tobacco <- 1
			t.paper <- 1
		}
		for _, smoker := range smokers {
			smoker <- next
		}
		wg.Add(1)
		wg.Wait()
	}
}

func smoker(t *Table, name string, smokes int, signal chan int, wg *sync.WaitGroup) {
	var chosen = -1
	for {
		chosen = <-signal // blocks

		if smokes != chosen {
			continue
		}

		fmt.Printf("Table: %d tobacco: %d match: %d\n", len(t.paper), len(t.tobacco), len(t.match))
		select {
		case <-t.paper:
		case <-t.tobacco:
		case <-t.match:
		}
		fmt.Printf("Table: %d tobacco: %d match: %d\n", len(t.paper), len(t.tobacco), len(t.match))
		time.Sleep(10 * time.Millisecond)
		select {
		case <-t.paper:
		case <-t.tobacco:
		case <-t.match:
		}
		fmt.Printf("Table: %d tobacco: %d match: %d\n", len(t.paper), len(t.tobacco), len(t.match))
		fmt.Printf("%s smokes a cigarette\n", name)
		time.Sleep(time.Millisecond * 500)
		wg.Done()
		time.Sleep(time.Millisecond * 100)
	}
}

const LIMIT = 1

func main() {
	var wg sync.WaitGroup
	table := new(Table)
	table.match = make(chan int, LIMIT)
	table.paper = make(chan int, LIMIT)
	table.tobacco = make(chan int, LIMIT)
	var signals [3]chan int
	// three smokers
	for i := 0; i < 3; i++ {
		signal := make(chan int, 1)
		signals[i] = signal
		go smoker(table, smokers[i], i, signal, &wg)
	}
	fmt.Printf("%s, %s, %s, sit with \n%s, %s, %s\n\n", smokers[0], smokers[1], smokers[2], smokeMap[0], smokeMap[1], smokeMap[2])
	arbitrate(table, signals, &wg)
}