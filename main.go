package main

import "fmt"

func main() {
	limit := getLimit()
	fmt.Println(limit())
	limit()
	fmt.Println(limit())
}

func getLimit() func() int {
	limit := 10
	return func() int {
		limit--
		return limit
	}
}
