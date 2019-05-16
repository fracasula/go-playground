package main

import "fmt"

// maxArea - O(n)
func maxArea(array []int) int {
	l := 0
	r := len(array) - 1
	max := 0
	for l <= r {
		h := array[l]
		if array[r] < h {
			h = array[r]
		}
		if area := h * (r - l); area > max { // r - l = distance
			max = area
		}
		if array[l] < array[r] {
			l++
		} else {
			r--
		}
	}

	return max
}

// maxArea2 - O(n²)
func maxArea2(array []int) int {
	l := 0
	r := len(array) - 1
	max := 0
	for l <= (r - 1) {
		for r >= (l + 1) {
			h := array[l]
			if array[r] < h {
				h = array[r]
			}
			distance := r - l
			area := h * distance
			if area > max {
				max = area
			}
			r--
		}
		l = r + 1
		r = len(array) - 1
	}

	return max
}

type scenario struct {
	input    []int
	expected int
}

var scenarios = []scenario{
	{
		input:    []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		expected: 49,
	},
	{
		input:    []int{1, 1},
		expected: 1,
	},
	{
		input:    []int{1, 10, 4, 5, 9, 3},
		expected: 27,
	},
}

func main() {
	for i, s := range scenarios {
		actual := maxArea(s.input)
		if s.expected != actual {
			fmt.Printf("Scenario %d failed (maxArea), got %d instead of %d", i+1, actual, s.expected)
		}
		actual = maxArea2(s.input)
		if s.expected != actual {
			fmt.Printf("Scenario %d failed (maxArea2), got %d instead of %d", i+1, actual, s.expected)
		}
	}
}
