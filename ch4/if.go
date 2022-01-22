package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 3
	if x, n := 4, rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
		fmt.Println(x)
	}
	fmt.Println(x)
}
