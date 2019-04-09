package main

import (
	"fmt"
	"os"
)

var pl = fmt.Println

func main() {
	pl("Reading Env")
	databasePass := os.Getenv("DATA_PASS")
	pl(databasePass)
}
