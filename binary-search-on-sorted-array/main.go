package main

import (
	"fmt"
)

func search(nums []int, num int) int {
	for l, r := 0, len(nums)-1; l <= r; {
		mid := int(l + (r-l)/2) // beware: cast to int is always a floor (not ceil)
		if nums[mid] == num {
			return mid
		} else if nums[mid] > num {
			r = mid - 1
		} else if nums[mid] < num {
			l = mid + 1
		}
	}

	return -1
}

func main() {
	scenarios := []scenario{
		{
			nums:     []int{1, 3, 6, 8, 9, 13, 16, 17, 18, 19, 20, 21, 22},
			num:      1,
			expected: 0,
		},
		{
			nums:     []int{1, 3, 6, 8, 9, 13, 16, 17, 18, 19, 20, 21, 22},
			num:      3,
			expected: 1,
		},
		{
			nums:     []int{1, 3, 6, 8, 9, 13, 16, 17, 18, 19, 20, 21, 22},
			num:      6,
			expected: 2,
		},
		{
			nums:     []int{1, 3, 6, 8, 9, 13, 16, 17, 18, 19, 20, 21, 22},
			num:      8,
			expected: 3,
		},
		{
			nums:     []int{1, 3, 6, 8, 9, 13, 16, 17, 18, 19, 20, 21, 22},
			num:      9,
			expected: 4,
		},
		{
			nums:     []int{1, 3, 6, 8, 9, 13, 16, 17, 18, 19, 20, 21, 22},
			num:      13,
			expected: 5,
		},
		{
			nums:     []int{1, 3, 6, 8, 9, 13, 16, 17, 18, 19, 20, 21, 22},
			num:      16,
			expected: 6,
		},
		{
			nums:     []int{1, 3, 6, 8, 9, 13, 16, 17, 18, 19, 20, 21, 22},
			num:      17,
			expected: 7,
		},
		{
			nums:     []int{1, 3, 6, 8, 9, 13, 16, 17, 18, 19, 20, 21, 22},
			num:      18,
			expected: 8,
		},
		{
			nums:     []int{1, 3, 6, 8, 9, 13, 16, 17, 18, 19, 20, 21, 22},
			num:      19,
			expected: 9,
		},
		{
			nums:     []int{1, 3, 6, 8, 9, 13, 16, 17, 18, 19, 20, 21, 22},
			num:      20,
			expected: 10,
		},
		{
			nums:     []int{1, 3, 6, 8, 9, 13, 16, 17, 18, 19, 20, 21, 22},
			num:      21,
			expected: 11,
		},
		{
			nums:     []int{1, 3, 6, 8, 9, 13, 16, 17, 18, 19, 20, 21, 22},
			num:      22,
			expected: 12,
		},
		{
			nums:     []int{1, 3, 6, 8, 9, 13, 16, 17, 18, 19, 20, 21, 22},
			num:      -999,
			expected: -1,
		},
		{
			nums:     []int{1, 3, 6, 8, 9, 13, 16, 17, 18, 19, 20, 21, 22},
			num:      999,
			expected: -1,
		},
	}

	for i, s := range scenarios {
		actual := search(s.nums, s.num)
		if actual != s.expected {
			fmt.Printf("Scenario %d failed, got %d instead of %d", i+1, actual, s.expected)
		}
	}
}

type scenario struct {
	nums     []int
	num      int
	expected int
}
