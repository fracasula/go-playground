package main

import (
	"fmt"
	"math"
)

func hourglassSum(arr [][]int32) int32 {
	var max int32 = math.MinInt32
	length := len(arr)
	for r := 0; r <= length-3; r++ {
		for c := 0; c <= length-3; c++ {
			var sum int32
			for j := c; j < c+3; j++ {
				sum += arr[r][j]
				sum += arr[r+2][j]
			}
			sum += arr[r+1][c+1]
			if sum > max {
				max = sum
			}
		}
	}

	return max
}

func main() {
	fmt.Println(
		hourglassSum([][]int32{
			{1, 1, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0},
			{0, 0, 2, 4, 4, 0},
			{0, 0, 0, 2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		}), // 19
	)
	fmt.Println(
		hourglassSum([][]int32{
			{-9, -9, -9, 1, 1, 1},
			{0, -9, 0, 4, 3, 2},
			{-9, -9, -9, 1, 2, 3},
			{0, 0, 8, 6, 6, 0},
			{0, 0, 0, -2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		}), // 28
	)
	fmt.Println(
		hourglassSum([][]int32{
			{-1, -1, 0, -9, -2, -2},
			{-2, -1, -6, -8, -2, -5},
			{-1, -1, -1, -2, -3, -4},
			{-1, -9, -2, -4, -4, -5},
			{-7, -3, -3, -2, -9, -9},
			{-1, -3, -1, -2, -4, -5},
		}), // -6
	)
}
