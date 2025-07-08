package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fizzbuzz(num int) string {
	var output []string
	if num%3 == 0 {
		output = append(output, "Fizz")
	}
	if num%5 == 0 {
		output = append(output, "Buzz")
	}
	if num%7 == 0 {
		output = append(output, "Bang")
	}
	if num%11 == 0 {
		output = []string{"Bong"}
	}
	if len(output) == 0 {
		output = []string{strconv.Itoa(num)}
	}
	return strings.Join(output, "")
}

// This is our main function, this executes by default when we run the main package.
func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz(i))
	}
}
