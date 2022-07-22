package main

import "fmt"

func Fizzbuzz1() {
	for i := 1; i <= 100; i++ {
		switch i % 15 {
		case 0:
			fmt.Println("FizzBuzz")
		case 3, 6, 9, 12:
			fmt.Println("Fizz")
		case 5, 10:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func Fizzbuzz2() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func main() {
	Fizzbuzz1()
	fmt.Println("----------------------")
	Fizzbuzz2()
}
