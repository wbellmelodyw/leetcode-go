package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var reverseNum int

	for x != 0 {
		last := x % 10
		x = x / 10
		reverseNum = reverseNum*10 + last
	}
	if reverseNum > math.MaxInt32 || reverseNum < math.MinInt32 {
		return 0
	}
	return reverseNum
}

func main() {
	fmt.Println(reverse(-1234))
}
