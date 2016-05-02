package main

import "fmt"

const (
	// NUM is a number of products
	NUM = 10
)

func producer(c chan<- int) {
	for i := 0; i < NUM; i++ {
		c <- i
	}
}

func consumer(c <-chan int, done chan<- bool) {
	for i := 0; i < NUM; i++ {
		val := <-c
		fmt.Println(val)
	}
	done <- true
}

func main() {
	c := make(chan int)
	done := make(chan bool)

	go producer(c)
	go consumer(c, done)

	<-done
}
