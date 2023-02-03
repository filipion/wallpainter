package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"example.com/wallpaint/cli"
	"example.com/wallpaint/square"
)

var PRICE_INSTRUCTION string = "Please specify the price of the paint you wish to use in EUR."

func main() {
	Form()
}

func Form() {
	fmt.Println("Welcome to the wall painter! Input \"quit\" at any time to quit the program")
	scanner := bufio.NewScanner(os.Stdin)
	area := 0.0
	wallList := []square.Square{}
	for {
		fmt.Printf(square.Instruct("wall"))
		wall, exitCode := square.ReadItem(scanner, "wall")
		if exitCode == "quit" {
			return
		}
		if exitCode == "end" {
			fmt.Println("Computing total wall area...")
			break
		}
		if exitCode == "cancel" {
			fmt.Println("Current wall was cancelled.")
			continue
		}
		wallList = append(wallList, wall)
	}
	fmt.Printf("\n=============================================\nLet's check that we got everything right. You inputed this data:\n")
	for {
		for i, wall := range wallList {
			fmt.Println("Id:", i, square.Represent(wall))
		}
		fmt.Println("Is there any wall you would like to change? If yes, specify which one by id. If not, press ENTER.")
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			break
		} else {
			id, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				fmt.Println("Error in reading input. Specify wall or press ENTER.")
			} else {
				fmt.Printf(square.Instruct("wall"))
				wall, exitCode := square.ReadItem(scanner, "wall")
				if exitCode == "quit" {
					return
				}
				if exitCode == "end" {
					fmt.Println("Cancelling wall rewrite.")
					break
				}
				if exitCode == "cancel" {
					fmt.Println("Cancelling wall rewrite.")
					continue
				}
				wallList[id] = wall
			}
		}
	}

	for _, wall := range wallList {
		area += square.Area(wall)
	}

	fmt.Printf("\n=============================================\nTotal wall area is %.2f m^2\n", area)

	for {
		price, exitCode := cli.InputValue(scanner, "price of paintjob per square meter (EUR)", "item", PRICE_INSTRUCTION)
		if exitCode == "" {
			fmt.Printf("Total price will be %.2f EUR", price*area)
			break
		} else if exitCode == "quit" {
			return
		} else {
			fmt.Printf("Please specify a price!")
		}
	}
}
