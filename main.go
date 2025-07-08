package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func reverse(list []string) []string {
	var res []string
	for i := len(list) - 1; i >= 0; i-- {
		res = append(res, list[i])
	}
	return res
}

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
	// Add Fezz immediately in front of the first thing beginning with B
	if num%13 == 0 {
		var fezzAdded = false
		for i, word := range output {
			if word[0] == 'B' {
				//temp := slices.(output[:i], "Fezz")
				//fmt.Println(temp)
				//fmt.Println(output)
				output = slices.Concat(output[:i], []string{"Fezz"}, output[i:])
				//fmt.Println(output)
				fezzAdded = true
				break
			}
		}
		if !fezzAdded {
			output = append(output, "Fezz")
			fezzAdded = true
		}
	}
	if num%17 == 0 {
		output = reverse(output)
	}
	if len(output) == 0 {
		output = []string{strconv.Itoa(num)}
	}
	return strings.Join(output, "")
}

// This is our main function, this executes by default when we run the main package.
func main() {
	for i := 11 * 13; i <= 11*13; i++ {
		fmt.Println(fizzbuzz(i))
	}
}
