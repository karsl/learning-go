package main

import (
	"errors"
	"fmt"
)

func divAndRemainder(numerator, denominator int) (result int, remainder int, _ error) {
	if denominator == 0 {
		return remainder, remainder, errors.New("cannot divide by zero")
	}
	result, remainder = numerator/denominator, numerator%denominator
	return
}

type person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func main() {
	fmt.Println(divAndRemainder(5, 2))

	p := person{
		name: "Test",
	}
	i := 2
	s := "Hello"
	modifyFails(i, s, p)
	fmt.Println(i, s, p)
}
