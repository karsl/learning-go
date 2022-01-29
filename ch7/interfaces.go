package main

import "fmt"

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string { // business logic
	fmt.Println("Running the logic")
	return data
}

func (lp LogicProvider) Process1(data string) string { // business logic
	fmt.Println("Running the logic1")
	return data
}

type Logic interface {
	Process(data string) string
}
type Client struct {
	L Logic
}

func (c Client) Program() {
	c.L.Process("data")
}

func main() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}
