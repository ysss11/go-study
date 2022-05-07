package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Fortune(val int) (result string) {

	switch val {
	case 10:
		result = "大吉"
	case 6, 7, 8, 9:
		result = "中吉"
	case 3, 4, 5:
		result = "吉"
	case 1, 2:
		result = "凶"
	case 0:
		result = "大凶"
	}

	return result
}

func main() {

	time := time.Now().UnixNano()
	rand.Seed(time)
	val := rand.Intn(11)

	fmt.Println(val)

	fmt.Println(Fortune(val))

}
