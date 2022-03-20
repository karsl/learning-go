package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		v := 1
		ch1 <- v
		fmt.Println("Reading value...")
		v2 := <-ch2
		fmt.Println("In goroutine: ", v, v2)
	}()

	v := func() int {
		fmt.Println("writing value v")
		return 2
	}
	var v2 int
	select {
	case ch2 <- v():
	case v2 = <-ch1:
	}

	fmt.Println("In main: ", v(), v2)

	time.Sleep(2 * time.Second)

	ch2 <- v()

	time.Sleep(2 * time.Second)

	fmt.Println("Let's try deadlock")

	select {
	case a := <-ch2:
		fmt.Println("Read from ch2 ", a)
	default:
		fmt.Println("I made you escape from deadlock!")
	}

	fmt.Println("You managed to escape from blocking select!")

	for {
		select {
		case a := <-ch2:
			fmt.Println("Read from ch2 ", a)
		default:
			fmt.Println("Default")
		}
	}
}
