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
	fmt.Println("Welcome to the wall painter!")
	scanner := bufio.NewScanner(os.Stdin)
	area := 0.0
	//lst := []int{}
	for {
		wall, exitCode := ReadItem(scanner, "wall")
		if exitCode == "quit" {
			break
		}
		if exitCode == "cancel" {
			fmt.Println("Current wall was cancelled.")
			continue
		}
		area += Area(wall)
	}
	fmt.Printf("\n\nTotal wall area is %.2f\n", area)

	for {
		price, exitCode := cli.InputValue(scanner, "price of paint (EUR)", "item", PRICE_INSTRUCTION)
		if exitCode == "" {
			fmt.Printf("Total price will be %.2f", price*area)
			break
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
	height, command := cli.InputValue(scanner, "height", item, instruct(item))
	if command != "" {
		return square{height: 0, width: 0}, command
	}
	width, command := cli.InputValue(scanner, "width", item, instruct(item))
	if command != "" {
		return square{height: 0, width: 0}, command
	}
	if item == "wall" {
		isDoubled, command = cli.InputBoolean(scanner, "Paint the wall on both sides?")
		if command != "" {
			return square{height: 0, width: 0}, command
		}

		hasOpenings, command := cli.InputBoolean(scanner, "Does the wall have doors, windows or other rectangular openings?")
		if command != "" {
			return square{height: 0, width: 0}, command
		}
		if hasOpenings {
			fmt.Println("Please specify the sizes of the windows/doors")
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
	return value, command
}

func instruct(item string) string {
	return fmt.Sprintf("Describe the %s in meters (for ex \"1.33\" or \"3\"). Type \"cancel\" to cancel this wall. Type \"quit\" to finish inputing walls:\n",
		item)
}
