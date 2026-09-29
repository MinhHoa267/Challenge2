package main

import "fmt"

func main() {
	fmt.Println("Tính tổng đến N:", tinhTong(5))
	demNguoc(5)
	fmt.Println("Số Fibonacci:", tinhFibonacci(5))
}
func tinhTong(n int) int {
	if n == 0 {
		return 0
	}
	return n + tinhTong(n-1)
}
func demNguoc(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Println("Đếm ngược:", n)
	return demNguoc(n - 1)
}

func tinhFibonacci(n int) int {
	if n == 0 {
		return 0
	}

	a := 0
	b := 1

	for i := 2; i <= n; i++ {
		tong := a + b
		fmt.Printf("Fibo: %d + %d = %d\n", a, b, tong)
		a = b
		b = tong
	}

	return b
}
