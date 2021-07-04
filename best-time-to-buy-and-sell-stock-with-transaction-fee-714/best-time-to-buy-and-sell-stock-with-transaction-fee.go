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

//
type student struct {
	Name string
	Age  int
}

func p() map[string]*student {
	m := make(map[string]*student)
	stus := []student{
		{Name: "z", Age: 1},
		{Name: "a", Age: 2},
		{Name: "b", Age: 3},
	}

	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	return m
}
func main() {
	//var a uint = 1
	//var b uint = 2
	//fmt.Println(a-b)
	//fmt.Println(uint(math.MaxUint64))//18446744073709551615
	//fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
	a := p()
	for k, v := range a {
		fmt.Printf("key=%s,value=%v \n", k, v)
	}
}
