package main

import (
	"fmt"
	"time"
)

func fizzbuzz(input int) chan int {
	result := make(chan int)
	go func() {
		for i := 1; i <= input; i++ {
			result <- i
		}
	}()
	return result
}

func main() {
	var input int
	fmt.Println("So luong phan tu ban muon: ")
	fmt.Scanln(&input)
	slice := []string{}
	now := time.Now()
	chan1 := fizzbuzz(input)
	for i := 1; i <= input; i++ {
		var ans = <-chan1
		var s = ""
		if ans%3 == 0 {
			s += "fizz"
		}
		if ans%5 == 0 {
			s += "buzz"
		}
		if (ans%3 != 0) && (ans%5 != 0) {
			s += fmt.Sprint(" ", ans)
		}
		slice = append(slice, s)
	}
	fmt.Printf("Result: %v", slice)
	fmt.Println(time.Since(now))
}
