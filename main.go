package main

import (
	"bufio"
	"fmt"
	"os"

	"example.com/wallpaint/cli"
)

var PRICE_INSTRUCTION string = "Please specify the price of the paint you wish to use in EUR."

type square struct {
	height    float64
	width     float64
	isDoubled bool
	unit      string
	windows   []square
}

func Area(wall square) float64 {
	area := wall.height * wall.width
	for _, window := range wall.windows {
		area -= window.height * window.width
	}
	if wall.isDoubled {
		area *= 2
	}
	return area
}

func main() {
	Form()
}

func Form() {
	fmt.Println("Welcome to the wall painter! Input `quit` at any time to quit. Input `end` to finish inputing a list of walls or windows.\nInput `cancel` to undo the current wall.")
	scanner := bufio.NewScanner(os.Stdin)
	area := 0.0
	for {
		fmt.Printf("\nDescribe a wall you wish to paint...")
		wall, exitCode := ReadItem(scanner, "wall")
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
		area += Area(wall)
	}
	fmt.Printf("\n=============================================\nTotal wall area is %.2f\n", area)

	for {
		price, exitCode := cli.InputValue(scanner, "price of paint (EUR)", "item", PRICE_INSTRUCTION)
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

func ReadItem(scanner *bufio.Scanner, item string) (square, string) {
	isDoubled := true
	windowList := []square{}
	fmt.Println()
	instruct(item)
	height, exitCode := cli.InputValue(scanner, "height", item, instruct(item))
	if exitCode != "" {
		return square{height: 0, width: 0}, exitCode
	}
	width, exitCode := cli.InputValue(scanner, "width", item, instruct(item))
	if exitCode != "" {
		return square{height: 0, width: 0}, exitCode
	}
	if item == "wall" {
		isDoubled, exitCode = cli.InputBoolean(scanner, "Paint the wall on both sides?")
		if exitCode != "" {
			return square{height: 0, width: 0}, exitCode
		}

		hasOpenings, command := cli.InputBoolean(scanner, "Does the wall have doors, windows or other rectangular openings?")
		if command != "" {
			return square{height: 0, width: 0}, command
		}
		if hasOpenings {
			fmt.Println("Please specify the sizes of the windows/doors. Type \"end\" after you are done.")
			for {
				window, exitCode := ReadItem(scanner, "window/door")
				if exitCode == "quit" {
					return square{}, exitCode
				}
				if exitCode == "end" {
					break
				}
				if exitCode == "cancel" {
					fmt.Println("Current window/door was cancelled.")
					continue
				}
				windowList = append(windowList, window)
			}
		}
	}

	value := square{height: height, width: width, isDoubled: isDoubled, windows: windowList}
	if item == "wall" {
		fmt.Println("Current wall results:", value)
	}
	return value, exitCode
}

func instruct(item string) string {
	return fmt.Sprintf("Describe the %s in meters (for ex \"1.33\" or \"3\"). Type \"cancel\" to cancel this wall. Type \"quit\" to finish inputing walls:\n",
		item)
}
