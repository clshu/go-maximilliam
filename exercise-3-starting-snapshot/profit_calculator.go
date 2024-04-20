package main

// Goals
// 1) Validate user input
//	=> Show error message & exit if invalid input is provided
// 	- No negative numbers
// 	- No 0
// 2) Store calculated results into file

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
		//  panic(err)
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	// if err1 != nil || err2 != nil || err3 != nil {
	// 	fmt.Println("No negative or zero")
	// 	return
	// }

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

	// writeNumbersToFile(revenue, expenses, taxRate)
	storeResults(ebt, profit, ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("The number should not be 0 or negative")
	}
	return userInput, nil
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
func writeNumbersToFile(ebt, profit, ratio float64) {

	// Create or open the file
	file, err := os.Create("calculator.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Create a buffered writer from the file
	writer := bufio.NewWriter(file)

	// List of lines to write
	ebtLine := fmt.Sprintf("Ebt: %0.1f", ebt)
	profitLine := fmt.Sprintf("Profit: %0.1f", profit)
	ratioLine := fmt.Sprintf("Ratio: %0.1f", ratio)
	lines := []string{ebtLine, profitLine, ratioLine}

	// Write lines to the buffer
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to buffer:", err)
			return
		}
	}

	// Flush the buffer to ensure all data is written to the file
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer:", err)
	}

}
