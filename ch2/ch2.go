package main

import "fmt"

const (
	key = "key"
	val = "val"
)

func main() {
	num1 := 1_0_01
	num1, num2 := 1_0_03, 2
	var num3, num4 = 3, 4
	num1 = 1_0_05
	var x int = 10
	var y float64 = 30.2
	var z float64 = float64(x) + y
	var d int = x + int(y)
	fmt.Println(num1 / num2)
	fmt.Println(z)
	fmt.Println(d)
	fmt.Println(key + " " + val)
	fmt.Println(num3 + num4)
}
