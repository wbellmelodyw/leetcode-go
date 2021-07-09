package main

import "fmt"

var m = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) (sum int) {
	n := len(s)
	for i := range s {
		cur := s[i]
		if i < n-1 && m[cur] < m[s[i+1]] {
			sum = sum - m[cur]
		} else {
			sum = sum + m[cur]
		}
	}
	return
}
func main() {
	fmt.Println(romanToInt("LVIII"))
}
