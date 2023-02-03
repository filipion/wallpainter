package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
	fmt.Printf("\n\nTotal wall area is %.2f", area)
}

func InputValue(scanner *bufio.Scanner, name string, item string, instructionsLine string) (float64, string) {
	var value float64
	fmt.Printf("Input %s:", name)

	for scanner.Scan() {
		if scanner.Text() == "cancel" || scanner.Text() == "quit" || scanner.Text() == "end" {
			break
		}

		input := scanner.Text()
		value, err := strconv.ParseFloat(input, 64)

		if err != nil {
			fmt.Println("Input Error!", instructionsLine)
			instruct(item)
		} else {
			return value, ""
		}
		fmt.Printf("Input %s:", name)
	}

	return value, scanner.Text()
}

func InputBoolean(scanner *bufio.Scanner, question string) (bool, string) {
	fmt.Printf("%s (y/n)", question)

	for scanner.Scan() {
		if scanner.Text() == "cancel" || scanner.Text() == "quit" {
			break
		}

		input := scanner.Text()

		if input != "y" && input != "n" {
			fmt.Println("Input Error!")
		} else if input == "y" {
			return true, ""
		} else {
			return false, ""
		}
		fmt.Printf(question)
	}

	return true, scanner.Text()
}

func ReadItem(scanner *bufio.Scanner, item string) (square, string) {
	isDoubled := true
	windowList := []square{}
	fmt.Println()
	instruct(item)
	height, command := InputValue(scanner, "height", item, instruct(item))
	if command != "" {
		return square{height: 0, width: 0}, command
	}
	width, command := InputValue(scanner, "width", item, instruct(item))
	if command != "" {
		return square{height: 0, width: 0}, command
	}
	if item == "wall" {
		isDoubled, command = InputBoolean(scanner, "Paint the wall on both sides?")
		if command != "" {
			return square{height: 0, width: 0}, command
		}

		hasOpenings, command := InputBoolean(scanner, "Does the wall have doors, windows or other rectangular openings?")
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
