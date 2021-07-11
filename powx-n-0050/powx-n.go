package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul2(x, n)
	}
	return 1.0 / quickMul2(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

//迭代法
func quickMul2(x float64, n int) float64 {
	ans := 1.0
	xContribute := x
	for n > 0 {
		if n%2 != 0 {
			ans *= xContribute
		}
		xContribute *= xContribute
		//右移
		n /= 2
	}
	return ans
}

func main() {
	fmt.Println(myPow(2, 6))
}
