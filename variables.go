package main

import (
	"fmt"
	"math/big"
	"strings"
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
}
