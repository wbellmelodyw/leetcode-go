package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	dpI0 := 0
	dpI1 := math.MinInt64
	for _, p := range prices {
		dpI0 = max(dpI0, dpI1+p)
		dpI1 = max(dpI1, dpI0-p-fee)
	}
	return dpI0
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
}
