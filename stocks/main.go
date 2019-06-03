package main

import "fmt"

func maxProfit(prices []int32) int32 {
	if len(prices)%2 != 0 {
		return 0
	}

	var min, max int32 = 0, 0
	for _, price := range prices {
		if price < min || min == 0 {
			min, max = price, 0
		} else if price > max {
			max = price
		}
	}

	if profit := max - min; profit < 0 {
		return 0
	} else {
		return profit
	}
}

func main() {
	fmt.Println(
		maxProfit([]int32{
			800, 854, 790, 823, 801, 803, 780, 801, 849, 790, 850, 830, // 70
		}),
		maxProfit([]int32{
			800, 854, // 54
		}),
		maxProfit([]int32{
			800, // 0
		}),
		maxProfit([]int32{
			1000, 900, // 0
		}),
		maxProfit([]int32{}),        // 0
		maxProfit([]int32{1, 2, 3}), // 0
	)
}
