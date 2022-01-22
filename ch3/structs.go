package main

import "fmt"

func main() {
	type person struct {
		name string
		age  int
		pet  string
	}

	p := person{
		age: 1,
	}

	p.name = "Sample name"

	fmt.Println(p)

	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "Bob",
		age:  50}
	var g struct {
		name string
		age  int
	}
	g.name = "Bob"
	g.age = 50

	fmt.Println(f == g)
}
