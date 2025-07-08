package main

import (
	"bufio"
	"fmt"
	"os"
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

func find(list []string, testFunc func(string) bool) int {
	for i, word := range list {
		if testFunc(word) {
			return i
		}
	}
	return len(list)
}

func getNum() int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the maximum number to print up to: ")
	scanner.Scan()
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid number, please try again")
		return getNum()
	}
	return num
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
		i := find(output, func(word string) bool { return word[0] == 'B' })
		output = slices.Concat(output[:i], []string{"Fezz"}, output[i:])
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
	maxnum := getNum()
	for i := 1; i <= maxnum; i++ {
		fmt.Println(fizzbuzz(i))
	}
}
