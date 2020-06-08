package main

import "fmt"

func checkSubarraySum(nums []int, k int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}
	m := make(map[int]int,0)
	m[0] = -1
	sum := 0
	for i:=0;i<length;i++{
		sum = sum+nums[i]
		if k != 0{
			sum = sum% k
		}
		if v,ok := m[sum];ok{
			if i-v >1{
				return true
			}
		}else{
			m[sum] = i
		}
	}
	return false
}

func main(){
	fmt.Println(checkSubarraySum([]int{2,3},0))
}
