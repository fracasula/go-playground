package main

import (
	"log"
	"math"
)

// https://leetcode.com/problems/string-to-integer-atoi/
func myAtoi(str string) int {
	l := len(str)
	if l == 0 {
		return 0
	}

	sign := 1
	var i int
	for ; i < l; i++ {
		if str[i] == ' ' {
			continue
		}
		if str[i] == '+' {
			i++
			break
		}
		if str[i] == '-' {
			sign = -1
			i++
			break
		}
		if str[i] < '0' || str[i] > '9' {
			return 0
		}
		break
	}

	res := 0
	for ; i < l; i++ {
		if str[i] < '0' || '9' < str[i] {
			break
		}
		res = res*10 + int(str[i]-'0')
		if res > math.MaxInt32 {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
	}

	return sign * res
}

func main() {
	scenarios := []scenario{
		{
			input:    "    -43 ciao",
			expected: -43,
		},
		{
			input:    "words and 987",
			expected: 0,
		},
		{
			input:    "42",
			expected: 42,
		},
		{
			input:    " 5",
			expected: 5,
		},
		{
			input:    "2147483648",
			expected: 2147483647,
		},
		{
			input:    " -91283472332",
			expected: -2147483648,
		},
		{
			input:    "9223372036854775808",
			expected: 2147483647,
		},
		{
			input:    "18446744073709551617",
			expected: 2147483647,
		},
		{
			input:    "  0000000000012345678",
			expected: 12345678,
		},
		{
			input:    "  -000000000000001",
			expected: -1,
		},
		{
			input:    "+-2",
			expected: 0,
		},
		{
			input:    "4193 with words",
			expected: 4193,
		},
		{
			input:    " +0 123",
			expected: 0,
		},
	}

	for i, s := range scenarios {
		actual := myAtoi(s.input)
		if actual != s.expected {
			log.Printf("Scenario %d failed, got %d instead of %d", i+1, actual, s.expected)
		}
	}
}

type scenario struct {
	input    string
	expected int
}
