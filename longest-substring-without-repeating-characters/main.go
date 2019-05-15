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
			fmt.Println("preIdx", preIdx, start)
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

//func lengthOfLongestSubstring(s string) int {
//	max := 0
//	seen := make(map[uint8]bool)
//
//	for i := 0; i < len(s); i++ {
//		_, ok := seen[s[i]]
//		if !ok {
//			seen[s[i]] = true
//		} else {
//			if max < len(seen) {
//				max = len(seen)
//			}
//			seen = make(map[uint8]bool)
//			i--
//		}
//	}
//
//	return max
//}

type scenario struct {
	input    string
	expected int
}

var scenarios = []scenario{
	{
		input:    "abcabcbb",
		expected: 3,
	},
	//{
	//	input:    "bbbbb",
	//	expected: 1,
	//},
	//{
	//	input:    "pwwkew",
	//	expected: 3,
	//},
	//{
	//	input:    "mamma",
	//	expected: 2,
	//},
	//{
	//	input:    "",
	//	expected: 0,
	//},
}

func main() {
	for i, s := range scenarios {
		actual := lengthOfLongestSubstring(s.input)
		if s.expected != actual {
			log.Printf("Scenario %d failed, got %d instead of %d", i+1, actual, s.expected)
		}
	}
}
