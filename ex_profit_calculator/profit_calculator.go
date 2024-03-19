package main

import "fmt"

func main() {
	var revenue float64
	var expense float64
	var taxRate float64
	var tax float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expense: ")
	fmt.Scan(&expense)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	earningBeforeTax := revenue - expense
	if earningBeforeTax > 0 {
		tax = earningBeforeTax * (taxRate / 100)
	}

	profit := earningBeforeTax - tax
	ratio := earningBeforeTax / profit

	fmt.Printf("Earning before Tax: %v\n", earningBeforeTax)
	fmt.Printf("Earning after Tax: %v\n", profit)
	fmt.Printf("Ratio: %v\n", ratio)
}
