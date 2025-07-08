package main

import "testing"

type Test struct {
	num      int
	expected string
}

var tests = []Test{
	{6, "Fizz"},
	{10, "Buzz"},
	{3 * 7, "FizzBang"},
	{3 * 11, "Bong"},
	{11 * 13, "FezzBong"},
	{3 * 5 * 17, "BuzzFizz"},
}

func TestFizzbuzz(t *testing.T) {
	for _, test := range tests {
		result := fizzbuzz(test.num)
		if result != test.expected {
			t.Errorf("incorrect output for num %d got: %s expected: %s", test.num, result, test.expected)
		}
	}
}
