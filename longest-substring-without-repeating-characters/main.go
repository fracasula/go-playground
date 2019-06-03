package main

import (
	"fmt"
	"log"
)

func lengthOfLongestSubstring(s string) int {
	left := 0
	start := 0
	seen := make(map[rune]int)
	for i, char := range s {
		fmt.Println(left)
		if preIdx, ok := seen[char]; ok && seen[char] >= start {
			start = preIdx + 1
			seen[char] = i
			continue
		}
		if distance := i - start + 1; distance > left {
			left = distance
		}
		seen[char] = i
	}
	return left
}

type scenario struct {
	input    string
	expected int
}

var scenarios = []scenario{
	{
		input:    "abcabcbb",
		expected: 3,
	},
	{
		input:    "bbbbb",
		expected: 1,
	},
	{
		input:    "pwwkew",
		expected: 3,
	},
	{
		input:    "mamma",
		expected: 2,
	},
	{
		input:    "",
		expected: 0,
	},
}

func main() {
	for i, s := range scenarios {
		actual := lengthOfLongestSubstring(s.input)
		if s.expected != actual {
			log.Printf("Scenario %d failed, got %d instead of %d", i+1, actual, s.expected)
		}
	}
}
