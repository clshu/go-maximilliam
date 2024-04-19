package main

import (
	"fmt"
)

func main() {
	var accountBalance float64 = 1000.0

	for {
		fmt.Println("\nWhat do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		default:
			fmt.Println("Invalid choice.")
		}

	}

}
