package main

import "testing"

type Test struct {
	num      int
	expected string
}

type TestWithOptions struct {
	num             int
	commandLineArgs []string
	expected        string
}

var tests = []Test{
	{6, "Fizz"},
	{10, "Buzz"},
	{3 * 7, "FizzBang"},
	{3 * 11, "Bong"},
	{11 * 13, "FezzBong"},
	{3 * 5 * 17, "BuzzFizz"},
}

var testsWithOptions = []TestWithOptions{
	{3 * 5 * 17, []string{"3"}, "Fizz"},
	{3 * 5 * 17, []string{"5"}, "Buzz"},
	{3 * 5 * 17, []string{"3", "5"}, "FizzBuzz"},
	{3 * 5 * 17, []string{"3", "5", "17"}, "BuzzFizz"},
}

func TestFizzbuzz(t *testing.T) {
	for _, test := range tests {
		result := fizzbuzz(test.num, []string{})
		if result != test.expected {
			t.Errorf("incorrect output for num %d got: %s expected: %s", test.num, result, test.expected)
		}
	}
}

func TestFizzbuzzWithOptions(t *testing.T) {
	for _, test := range testsWithOptions {
		result := fizzbuzz(test.num, test.commandLineArgs)
		if result != test.expected {
			t.Errorf("incorrect output for num %d got: %s expected: %s", test.num, result, test.expected)
		}
	}
}
