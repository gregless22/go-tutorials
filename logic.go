package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// // var x float64 = 42
	// var result string

	// if x := -2; x < 0 {
	// 	result = "Less than Zero"
	// 	fmt.Println(result)
	// }

	// result = "Greater than or Eqaul to Zero"

	// fmt.Println(result)

	// Swithc
	rand.Seed(time.Now().Unix())
	// dow := rand.Intn(6) + 1
	// fmt.Println("day", dow)

	var result string

	switch dow := rand.Intn(6) + 1; dow {
	case 1:
		result = "It's Sunday"
	case 2:
		result = "It's Saturady"
	default:
		result = "It's a weekday"
	}
	fmt.Println(result)

	// replaceing if else
	x := 0
	switch {
	case x < 0:
		result = "less than zero"
	case x == 0:
		result = "equal to zero"
	default:
		result = "Greater than zero"
	}
	fmt.Println(result)

	// while
	sum := 1
	fmt.Println("SUM: ", sum)

	colors := []string{"red", "green", "blue"}
	fmt.Println(colors)

	sum = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("SUM: ", sum)

	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println("SUM: ", sum)
	}
}
