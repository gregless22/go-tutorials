package main

import (
	"fmt"
	"sync"
)

var (
	mutex sync.Mutex
	pf    = fmt.Printf
	pl    = fmt.Println
)

func main() {

}
