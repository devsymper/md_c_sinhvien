package main

import (
	"fmt"
	"time"
)

func printFoo(a int) chan string {
	chanFoo := make(chan string)
	go func() {
		for i := 0; i < a; i++ {
			chanFoo <- "foo"
		}
	}()
	return chanFoo
}

func printBar(a int) chan string {
	chanBar := make(chan string)
	go func() {
		for i := 0; i < a; i++ {
			chanBar <- "bar"
		}
	}()
	return chanBar
}

func main() {
	now := time.Now()
	var s string
	var input int
	fmt.Println("So lan ban muon in foobar: ")
	fmt.Scanln(&input)
	foo := printFoo(input)
	bar := printBar(input)
	for i := 0; i < input; i++ {
		s += <-foo + <-bar
	}
	fmt.Println(s)
	fmt.Println(time.Since(now))
}
