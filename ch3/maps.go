package main

import "fmt"

func main() {
	var nilMap map[string]int
	fmt.Println(nilMap == nil)
	fmt.Println(nilMap["asd"])
	emptyMap := map[string]int{}
	fmt.Println(emptyMap == nil)
	fmt.Println(emptyMap["asd"])
	fmt.Println(len(emptyMap))
}
