package main

import "fmt"

func main() {
	var revenue float64
	var expense float64
	var taxRate float64

	revenue = acceptInput("Revenue: ")
	expense = acceptInput("Expense: ")
	taxRate = acceptInput("Tax Rate: ")
	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)
	// fmt.Print("Expense: ")
	// fmt.Scan(&expense)
	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)

	// earningBeforeTax := revenue - expense
	// if earningBeforeTax > 0 {
	// 	tax = earningBeforeTax * (taxRate / 100)
	// }

	// profit := earningBeforeTax - tax
	// ratio := earningBeforeTax / profit

	earningBeforeTax, profit, ratio := calculateFinancial(revenue, expense, taxRate)

	fmt.Printf("Earning before Tax: %.2f\n", earningBeforeTax)
	fmt.Printf("Earning after Tax: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func acceptInput(message string) float64 {
	var input float64
	fmt.Print(message)
	fmt.Scan(&input)

	return input
}

func calculateFinancial(revenue, expense, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expense
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return
}
