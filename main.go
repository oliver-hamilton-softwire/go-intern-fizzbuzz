package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(num int) string {
	output := ""
	if num%3 == 0 {
		output += "Fizz"
	}
	if num%5 == 0 {
		output += "Buzz"
	}
	if num%7 == 0 {
		output += "Bang"
	}
	if num%11 == 0 {
		output = "Bong"
	}
	if output == "" {
		output = strconv.Itoa(num)
	}
	return output
}

// This is our main function, this executes by default when we run the main package.
func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz(i))
	}
}
