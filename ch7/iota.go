package main

import "fmt"

type MailCategory int

const (
	Uncategorized MailCategory = iota + 1
	_
	Personal
	Spam
	Social
	Advertisements
)

func main() {
	fmt.Println(Advertisements)
}
