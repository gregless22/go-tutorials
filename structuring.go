package main

import "fmt"

// Dog comment
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

// Animal comment
type Animal interface {
	Speak()
}

// Speak comment
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

// SpeakThreeTimes Comment
func (d *Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v! %v! %v!", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

func main() {
	poodle := Dog{"Poodle", 23, "Woof"}
	fmt.Println(poodle)
	poodle.Speak()

	poodle.Sound = "arrf"
	poodle.Speak()

	poodle.SpeakThreeTimes()
	poodle.SpeakThreeTimes()
	// poodle.SpeakThreeTimes()

	lab := Animal(Dog{"Labrador", 45, "Woof Woof_"})
	lab.Speak()
}
