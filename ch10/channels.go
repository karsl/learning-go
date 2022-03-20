package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(10 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(15 * time.Second)
		ch2 <- 2
	}()

	select {
	case a := <-ch1:
		fmt.Println("Received this from ch1: ", a)
	case b := <-ch2:
		fmt.Println("Received this from ch2: ", b)
	}
}
