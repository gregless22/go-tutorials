package main

import (
	"fmt"
	"math/big"
	"strings"
	"time"
)

var (
	pl = fmt.Println
)

func main() {
	str1 := "An implicitly typed string"
	fmt.Printf("str1: %v: %T\n", str1, str1)

	str2 := "An explicitly typed string"
	fmt.Printf("str1: %v: %T\n", str2, str2)

	fmt.Println(strings.ToUpper(str1))
	pl(strings.Title(str1))

	var b1, b2, b3, bigSum big.Float

	b1.SetFloat64(23.5)
	b2.SetFloat64(65.1)
	b3.SetFloat64(76.3)

	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	fmt.Printf("BigSum = %.10g\n", &bigSum)

	// add some data stuff
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go Launchdate at %s\n", t)
	now := time.Now()

	fmt.Printf("The time now is %s\n", now)

	fmt.Println("the month is ", t.Month())
	fmt.Println("the day is ", t.Day())
	fmt.Println("the weekday is ", t.Weekday())

	tomorrow := t.AddDate(0, 0, 1)
	fmt.Println("the month is ", tomorrow.Month())
	fmt.Println("the day is ", tomorrow.Day())
	fmt.Println("the weekday is ", tomorrow.Weekday())

}
