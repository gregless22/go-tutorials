package main

import "fmt"

var (
	pl = fmt.Println
)

func main() {
	pl("Hello World")
}

// Calculate takes an int and returns doubl
func Calculate(x int) (int, error) {
	result := x + 2
	return result, nil
}
