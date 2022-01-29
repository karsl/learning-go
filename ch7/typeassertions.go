package main

import (
	"fmt"
	"os"
)

type MyInt int

func main() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	i2, ok := i.(int)
	if !ok {
		fmt.Fprintf(os.Stderr, "unexpected type for %v\n", i)
	} else {
		fmt.Println(i2)
	}

	switch i.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int")
	case MyInt:
		fmt.Println("MyInt")
	default:
		fmt.Println("otherwise")
	}
}
