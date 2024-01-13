package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
		// panic("Can't progress")
	}

	fmt.Println("Welcome to Go Bank!")

	for {

		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your current balance is:", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance Updated! New Amount: ", accountBalance)
		case 3:
			fmt.Print("Your withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have")
				continue
			}
			accountBalance -= withdrawAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance Updated! New Amount: ", accountBalance)

		default:
			fmt.Println("Thank you for using our service. Catch you later!")
			return

			// if choice == 1 {
			// 	fmt.Println("Your current balance is:", accountBalance)
			// } else if choice == 2 {
			// 	fmt.Print("Your deposit: ")
			// 	var depositAmount float64
			// 	fmt.Scan(&depositAmount)
			// 	if depositAmount <= 0 {
			// 		fmt.Println("Invalid amount. Must be greater than 0.")
			// 		continue
			// 	}
			// 	accountBalance += depositAmount
			// 	fmt.Println("Balance Updated! New Amount: ", accountBalance)
			// } else if choice == 3 {
			// 	fmt.Print("Your withdraw: ")
			// 	var withdrawAmount float64
			// 	fmt.Scan(&withdrawAmount)
			// 	if withdrawAmount <= 0 {
			// 		fmt.Println("Invalid amount. Must be greater than 0.")
			// 		return
			// 	}

			// 	if withdrawAmount > accountBalance {
			// 		fmt.Println("Invalid amount. You can't withdraw more than you have")
			// 	}
			// 	accountBalance -= withdrawAmount
			// 	fmt.Println("Balance Updated! New Amount: ", accountBalance)
			// } else {
			// 	fmt.Println("Thank you for using our service. Catch you later!")
			// 	break
			// }
			//
		}
	}
}
