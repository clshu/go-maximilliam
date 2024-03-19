package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {

	var investmenAmount float64
	expectedReturnRate := 5.5
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmenAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Investment Years: ")
	fmt.Scan(&years)

	// futureValue := investmenAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	futureValue, futureRealValue := calculateFutureValues(investmenAmount, expectedReturnRate, years)
	// fmt.Println("Future Value: ", futureValue)
	// fmt.Println("Future Value(adjusted for inflation): ", futureRealValue)
	// fmt.Printf("Future Value: %.2f\n", futureValue)
	// fmt.Printf("Future Value(adjusted for inflation): %.2f\n", futureRealValue)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	fotmattedRFV := fmt.Sprintf("Future Value(adjusted for inflation): %.2f\n", futureRealValue)

	fmt.Print(formattedFV, fotmattedRFV)
}

func calculateFutureValues(investmenAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmenAmount * math.Pow(1+expectedReturnRate/100, years)
	frv := fv / math.Pow(1+inflationRate/100, years)

	return fv, frv
}
