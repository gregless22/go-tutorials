package main

import "fmt"

// TODO need to lookat go modules

func main() {
	doSomething()
	sum := addValues(22, 54)
	fmt.Println(sum)

	sum = addAllValues(12, 54, 79)
	fmt.Println(sum)

	n1, l1 := stringutil.FullName("greg", "connolly")
	fmt.Printf("Fullname: %v, number of chars %v\n", n1, l1)

	n2, l2 := stringutil.FullNameNakedReturn("fdsagafdg", "gfdsg")
	fmt.Printf("Fullname: %v, number of chars %v\n", n2, l2)
}

func doSomething() {
	fmt.Println("Doing Something")
}

func addValues(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	fmt.Printf("%T\n", values)
	return sum
}
