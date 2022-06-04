package main

import (
	"fmt"
	"math"
)

func Fibonacci_01() func() int {
	prev := 0
	fibo := -1

	return func() int {
		temp := fibo
		if fibo == -1 {
			fibo = 0
			temp = 0
		} else if fibo == 0 {
			fibo += 1
		} else {
			fibo += prev
		}
		prev = temp
		return fibo
	}
}
func Fibonacci_02() func() int {
	n, n1 := 0, 1
	return func() int {
		v := n
		n, n1 = n1, n+n1
		return v
	}
}

func Fibonacci_03(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci_03(n-2) + Fibonacci_03(n-1)
}

func Fibonacci_04(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func Fibonacci_05(n int, memo map[int]int) int {
	if n < 2 {
		return n
	}
	if _, ok := memo[n]; !ok {
		memo[n] = Fibonacci_05(n-2, memo) + Fibonacci_05(n-1, memo)
	}
	return memo[n]
}

func Fibonacci_06(n int) float64 {
	return math.Round((math.Pow((1+math.Sqrt(5))/2, float64(n)) - math.Pow((1-math.Sqrt(5))/2, float64(n))) / math.Sqrt(5))
}

// フィボナッチ数を表示するコード
func main() {

	f01 := Fibonacci_01()
	f02 := Fibonacci_02()

	for i := 0; i < 10; i++ {
		fmt.Print(f01())
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Print(f02())
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Print(Fibonacci_03(i))
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Print(Fibonacci_04(i))
	}
	fmt.Println()
	memo := make(map[int]int)
	for i := 0; i < 10; i++ {
		fmt.Print(Fibonacci_05(i, memo))
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Print(Fibonacci_06(i))
	}
	fmt.Println()
}
