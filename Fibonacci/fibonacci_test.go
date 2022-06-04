package main

import "testing"

const maxCount = 10000

func BenchmarkFibonacci_01(b *testing.B) {
	f := Fibonacci_01()
	for i := 0; i < maxCount; i++ {
		f()
	}
}

func BenchmarkFibonacci_02(b *testing.B) {
	f := Fibonacci_02()
	for i := 0; i < maxCount; i++ {
		f()
	}
}

// 失敗するためコメントアウト
// func BenchmarkFibonacci_03(b *testing.B) {
// 	for i := 0; i < maxCount; i++ {
// 		Fibonacci_03(i)
// 	}
// }

func BenchmarkFibonacci_04(b *testing.B) {
	for i := 0; i < maxCount; i++ {
		Fibonacci_04(i)
	}
}

func BenchmarkFibonacci_05(b *testing.B) {
	memo := make(map[int]int)
	for i := 0; i < maxCount; i++ {
		Fibonacci_05(i, memo)
	}
}

func BenchmarkFibonacci_06(b *testing.B) {
	for i := 0; i < maxCount; i++ {
		Fibonacci_06(i)
	}
}
