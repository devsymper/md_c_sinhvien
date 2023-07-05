package main

import (
	"fmt"
	"time"
)

func fizz() chan string {
	result := make(chan string)
	go func() {
		result <- "fizz"
	}()
	return result
}

func buzz() chan string {
	result := make(chan string)
	go func() {
		result <- "buzz"
	}()
	return result
}

func main() {
	var input int
	fmt.Println("So luong phan tu ban muon: ")
	fmt.Scanln(&input)
	slice := []string{}
	now := time.Now()
	chan1 := fizz()
	chan2 := buzz()
	for i := 1; i <= input; i++ {
		var s = ""
		if i%3 == 0 {
			s += <-chan1
		}
		if i%5 == 0 {
			s += <-chan2
		} else {
			s += fmt.Sprintf("%v", i)
		}
		slice = append(slice, s)
	}
	fmt.Printf("Result: %v", slice)
	fmt.Println(time.Since(now))
}
