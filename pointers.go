package main

import "fmt"

// Dog sdf
type Dog struct {
	Breed  string
	Weight int
}

func main() {
	var p *int
	if p != nil {
		fmt.Println("Value of P: ", *p)
	}

	// Arrays
	var colors [3]string
	colors[0] = "red"
	colors[1] = "blue"
	colors[2] = "green"
	fmt.Println(colors)

	var numbers = [5]int{5, 3, 1, 2, 4}
	fmt.Println(numbers)
	fmt.Println("number of colors: ", len(colors))

	// Slices
	var colours = []string{"red", "green", "blue"}
	fmt.Println(colours)

	colours = append(colours, "purple")
	fmt.Println(colours)

	colours = append(colours[1:len(colours)])
	fmt.Println(colours)

	// Maps

	states := make(map[string]string)
	fmt.Println(states)

	// add some items
	states["WA"] = "washington"
	states["OR"] = "Oregan"
	states["CA"] = "California"

	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	poodle := Dog{"poodle", 34}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)

}
