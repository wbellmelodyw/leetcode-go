package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	dp := make([][2]int, n)
	for i, p := range prices {
		if i-1 == -1 {
			dp[0][0] = 0
			dp[0][1] = -p
		} else {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+p)
			if i-2 < 0 {
				dp[i][1] = max(dp[i-1][1], -p)
			} else {
				dp[i][1] = max(dp[i-1][1], dp[i-2][0]-p)
			}
		}
	}
	return dp[n-1][0]
}

func maxProfit2(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	dpI0 := 0
	dpI1 := math.MinInt64
	dpPre := 0 //前一天
	for _, p := range prices {
		temp := dpI0
		dpI0 = max(dpI0, dpI1+p)
		dpI1 = max(dpI1, dpPre-p)
		dpPre = temp
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
	fmt.Println(maxProfit2([]int{1, 2, 3, 0, 2}))
}
