package main

import (
	"fmt"
)

//给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
//设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。

func maxProfit(prices []int) int {
	//最大交易数
	const maxK = 2
	length := len(prices)
	dp := make([][maxK + 1][2]int, length)
	for i := 0; i < length; i++ {
		if i-1 == -1 {
			//初始化
			dp[0][1][0] = 0
			dp[0][2][0] = 0
			dp[0][1][1] = -prices[i]
			dp[0][2][1] = -prices[i]
			continue
		}
		for k := maxK; k >= 1; k-- {
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[length-1][maxK][0]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
