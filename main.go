package main

import (
	"fmt"
	"time"
)

var (
	pf = fmt.Printf
	pl = fmt.Println
)

func main() {
	coolFunc(myFunc)
}

func myFunc() {
	pl("Hello World")
	time.Sleep(1 * time.Second)
}

func coolFunc(a func()) {
	pf("Starting function execution: %s\n", time.Now())
	a()
	pf("End of function execution: %s\n", time.Now())
}
