
package main

import "fmt"

type Producer struct{
	msg *chan int
	done *chan bool
}

type Consumer struct{
	msg *chan int
}

func NewProducer(msg *chan int, done *chan bool) *Producer{
	return &Producer{msg, done}
}

func NewConsumer(msg *chan int) *Consumer{
	return &Consumer{msg}
}

func (p *Producer) produce(max int){
	fmt.Println("Producer Started")
	for i:= 0; i<max; i++ {
		fmt.Println("Produced", i+1)
		*p.msg <- i+1
	}
	fmt.Println("Producer done")
	*p.done <- true
}

func (c *Consumer) consume() {
	fmt.Println("Consumer Started")
	for true{
		 msg := <-*c.msg
		fmt.Println("Consumed", msg)
	}
	//fmt.Println("Consumer finished")
}

func main(){
	msg := make(chan int)
	done := make(chan bool)

	go NewProducer(&msg, &done).produce(20)
	go NewConsumer(&msg).consume()
	<- done
}
