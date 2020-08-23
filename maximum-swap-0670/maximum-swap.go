package main

import (
	"fmt"
	"strconv"
)

func maximumSwap(num int) int {
	numBytes := []byte(strconv.Itoa(num))
	data := []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	for i, numByte := range numBytes {
		t := int(numByte - '0')
		data[t] = i
	}

	for ii, numB := range numBytes {
		now := int(numB - '0')
		for j := 9; j > now; j-- {
			if data[j] > ii {
				numBytes[ii], numBytes[data[j]] = numBytes[data[j]], numBytes[ii]
				q, _ := strconv.Atoi(string(numBytes))
				return q
			}
		}
	}
	return num
}

func main() {
	fmt.Println(maximumSwap(98368))
}
