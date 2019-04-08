package main

import "fmt"

var (
	pl   = fmt.Println
	name string
)

func init() {
	pl("this is called ont he main init")
	name = "Greg"
}

func main() {
	pl("this is called after")
	pl(name)
}
