package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		comment := strconv.Itoa(i)
		if i%3 == 0 {
			comment = "Fizz"
		} else if i%5 == 0 {
			comment = "Buzz"
		}
		if i%3 == 0 && i%5 == 0 {
			comment = "FizzBuzz"
		}
		fmt.Printf("%s\n", comment)
	}
}
