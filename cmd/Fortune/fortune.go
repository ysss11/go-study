package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	time := time.Now().UnixNano()
	rand.Seed(time)
	val := rand.Intn(11)

	fmt.Println(val)

	switch val {
	case 10:
		fmt.Println("大吉")
	case 6, 7, 8, 9:
		fmt.Println("中吉")
	case 3, 4, 5:
		fmt.Println("吉")
	case 1, 2:
		fmt.Println("凶")
	case 0:
		fmt.Println("大凶")
	}

}
