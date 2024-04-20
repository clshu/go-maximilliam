package main

import "fmt"

func main() {
	age := 32

	var agePointer *int

	agePointer = &age

	fmt.Println("Age:", *agePointer)
	fmt.Println("Age Pinter:", agePointer)
	getAdultYear(agePointer)
	fmt.Println("Adult Year:", age)
	fmt.Println("Age:", age)
}

func getAdultYear(age *int) {
	// return *age - 18
	*age = *age - 18
}
