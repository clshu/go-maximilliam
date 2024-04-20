package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const fileName string = "balance.txt"
const defaultValue float64 = 1000.0

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(fileName, defaultValue)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println("-------------")
		fmt.Println(err)
		// panic(err)
	}

	for {
		displayOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Balance:", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			accountBalance += depositAmount
			fmt.Println("Balance updated! New balance: ", accountBalance)
			fileops.WriteFloatToFile(fileName, accountBalance)
		case 3:
			fmt.Print("Withdrawal Amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if accountBalance < withdrawalAmount {
				fmt.Println("Not sufficient fund. Current balance: ", accountBalance)
				continue
			}
			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New balance: ", accountBalance)
			fileops.WriteFloatToFile(fileName, accountBalance)
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		default:
			fmt.Println("Invalid choice.")
		}

	}
}
