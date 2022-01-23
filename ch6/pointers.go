package main

import "fmt"

func main() {
	f2()
}

func f1() {
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX) // prints a memory address fmt.Println(*pointerToX) // prints 10
	z := 5 + *pointerToX
	fmt.Println(z) // prints 15
}

func f2() {
	type Foo struct {
	}
	x := &Foo{}
	var y int
	z := &y
	fmt.Println("x:", x, "y:", y, "z:", z)
}
