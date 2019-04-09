package main

import "fmt"

func main() {
	myVariadicFunctions("Heloo", "world")
}

func myVariadicFunctions(args ...string) {
	fmt.Println(args)
}
