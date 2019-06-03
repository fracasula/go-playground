package main

import (
	"log"
	"reflect"
)

// https://leetcode.com/problems/rotate-array/
func rotate(nums []int, k int) {
	l := len(nums)
	rotated := make([]int, l)
	for i := 0; i < l; i++ {
		rotated[(i+k)%l] = nums[i]
	}
	copy(nums, rotated)
}

func main() {
	scenarios := []scenario{
		{
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums:     []int{1, 2, 3, 4, 5},
			k:        5,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			nums:     []int{1, 2, 3, 4, 5},
			k:        1,
			expected: []int{5, 1, 2, 3, 4},
		},
		{
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        1,
			expected: []int{7, 1, 2, 3, 4, 5, 6},
		},
	}

	for i, s := range scenarios {
		rotate(s.nums, s.k)
		if !reflect.DeepEqual(s.nums, s.expected) {
			log.Printf("Scenario %d failed, got %v instead of %v", i+1, s.nums, s.expected)
		}
	}
}

type scenario struct {
	nums     []int
	k        int
	expected []int
}
