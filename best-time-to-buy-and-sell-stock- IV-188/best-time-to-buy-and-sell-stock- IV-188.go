package main

import (
	"fmt"
	"math"
)

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if k > (n / 2) {
		//k相当于无穷大
		return greedy(prices)
	}
	dp := make([][][2]int, n)
	for i, p := range prices {
		dp[i] = make([][2]int, k+1)
		for j := 1; j <= k; j++ {
			if i-1 == -1 { //base case
				dp[0][j][0] = 0
				dp[0][0][1] = math.MaxInt64
				dp[0][j][1] = -p
			} else {
				dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+p)
				dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-p)
			}
		}
	}
	return dp[n-1][k][0]
}
func greedy(prices []int) int {
	sum := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			sum += prices[i] - prices[i-1]
		}
	}
	return sum
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}
