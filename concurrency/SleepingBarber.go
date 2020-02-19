////The problem describes two processes, the producer and the consumer, who share a common, fixed-size buffer used as a
////queue. The producer's job is to generate data, put it into the buffer, and start again. At the same time, the consumer
////is consuming the data (i.e., removing it from the buffer), one piece at a time. The problem is to make sure that the
////producer won't try to add data into the buffer if it's full and that the consumer won't try to remove data from an empty
////buffer. The solution for the producer is to either go to sleep or discard data if the buffer is full. The next time the
////consumer removes an item from the buffer, it notifies the producer, who starts to fill the buffer again. In the same
////way, the consumer can go to sleep if it finds the buffer empty. The next time the producer puts data into the buffer,
////it wakes up the sleeping consumer.
//
package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	sleeping = iota
	checking
	cutting
)
const sTime = time.Millisecond * 100

var states = []string{"Sleeping", "Checking", "Cutting"}


type Barber struct {
	name string
	sync.Mutex
	state    int
	customer *Customer
}

type Customer struct {
	name string
}

func (c *Customer) String() string {
	return c.name
}
func NewCustomer(name string) *Customer{
	return &Customer{name: name}
}

func NewBarber(name string) *Barber {
	return &Barber{name:  name, state: sleeping}
}


func barber(b *Barber, wr chan *Customer, wakers chan *Customer, wg *sync.WaitGroup) {
	for {
		b.Lock()
		b.state = checking
		b.customer = nil
		fmt.Printf("Checking waiting room: %d\n", len(wr))
		time.Sleep(sTime)
		select {
		case c := <-wr:
			HairCut(c, b, wg)
			b.Unlock()
		default:
			fmt.Printf("Sleeping Barber - %s\n", b.customer)
			b.state = sleeping
			b.customer = nil
			b.Unlock()
			c := <-wakers
			b.Lock()
			fmt.Printf("Woken by %s\n", c)
			HairCut(c, b, wg)
			b.Unlock()
		}
	}
}

func HairCut(c *Customer, b *Barber, wg *sync.WaitGroup) {
	b.state = cutting
	b.customer = c
	b.Unlock()
	fmt.Printf("Cutting  %s hair\n", c)
	time.Sleep(sTime)
	b.Lock()
	wg.Done()
	b.customer = nil
}


func customer(c *Customer, b *Barber, wr chan *Customer, wk chan *Customer, wg *sync.WaitGroup) {
	time.Sleep(sTime)
	b.Lock()
	fmt.Printf("Customer %s checks %s barber | room: %d, w %d - customer: %s\n",
		c, states[b.state], len(wr), len(wk), b.customer)
	switch b.state {
	case sleeping:
		select {
		case wk <- c:
		default:
			select {
			case wr <- c:
			default:
				wg.Done()
			}
		}
	case cutting:
		select {
		case wr <- c:
		default:
			wg.Done()
		}
	case checking:
		panic("Impossible Situation")
	}
	b.Unlock()
}

func main() {
	b := NewBarber("Raj")
	waitRoom := make(chan *Customer, 6)
	wakers := make(chan *Customer, 1)

	nCustomers := 10
	var wg sync.WaitGroup
	wg.Add(nCustomers)

	go barber(b, waitRoom, wakers, &wg)
	time.Sleep(sTime)


	for i := 0; i < nCustomers; i++ {
		time.Sleep(sTime)
		c := NewCustomer(fmt.Sprintf("Customer-%d", i+1))
		go customer(c, b, waitRoom, wakers, &wg)
	}

	wg.Wait()
	fmt.Println("Done for the day")
}