package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for j, num := range nums {
		if index, ok := m[target-num]; ok {
			return []int{index, j}
		} else {
			m[num] = j
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 5}, 9))
}
