package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	res := 0
	i_2 := 1
	i_1 := 2
	for i := 2; i < n; i++ {
		res = i_1 + i_2
		i_2 = i_1
		i_1 = res
	}
	return res
	//	if n==1{
	//		return 1
	//	}
	//	res := make([]int,n)
	//	res[0] = 1
	//	res[1] = 2
	//	for i:=2;i<n;i++{
	//		res[i] = res[i-1]+res[i-2]
	//	}
	//	return res[n-1]
}
func main() {
	fmt.Println(climbStairs(4))
}
