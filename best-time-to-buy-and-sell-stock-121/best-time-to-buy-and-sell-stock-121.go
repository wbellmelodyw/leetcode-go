package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	dp_i_0 := 0
	dp_i_1 := math.MinInt64
	for _, p := range prices {
		dp_i_0 = max(dp_i_0, dp_i_1+p)
		dp_i_1 = max(dp_i_1, -p)
	}
	return dp_i_0
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
