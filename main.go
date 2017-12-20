package main

import "fmt"

func sum(arr []int, ch chan int) {
	sum := 0

	for _, v := range arr {
		sum += v
	}

	ch <- sum
}

func main() {
	arr := []int{11, 22, 33, 44, 55}
	ch := make(chan int)

	go sum(arr, ch)

	fmt.Println(arr, <-ch)
}
