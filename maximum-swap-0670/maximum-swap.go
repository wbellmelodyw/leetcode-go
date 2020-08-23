package main

import (
	"fmt"
	"strconv"
)

func maximumSwap(num int) int {
	numBytes := []byte(strconv.Itoa(num))
	data := make(map[int]int)
	for i, numByte := range numBytes {
		t := int(numByte - '0')
		data[t] = i
	}

	for ii, numB := range numBytes {
		for j := 9; j > int(numB-'0'); j-- {
			if index, ok := data[j]; ok {
				if index > ii {
					numBytes[ii], numBytes[index] = numBytes[index], numBytes[ii]
					q, _ := strconv.Atoi(string(numBytes))
					return q
				}
			}
		}
	}
	return num
}

func main() {
	fmt.Println(maximumSwap(98368))
}
