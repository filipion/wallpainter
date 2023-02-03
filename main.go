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
}

func main() {
	Form()
}

func Form() {
	fmt.Println("Welcome to the wall painter! Press `q` to quit.")
	scanner := bufio.NewScanner(os.Stdin)
	area := 0.0
	//lst := []int{}
	for {
		wall, exitCode := ReadWall(scanner)
		if exitCode == "quit" {
			break
		}
		if exitCode == "cancel" {
			fmt.Println("Current wall was cancelled.")
			continue
		}
		area += wall.height * wall.width
	}
	fmt.Printf("\n\nTotal wall area is %.2f", area)
}

func InputValue(scanner *bufio.Scanner, name string, item string) (float64, string) {
	var value float64
	fmt.Printf("Input %s:", name)

	for scanner.Scan() {
		if scanner.Text() == "cancel" || scanner.Text() == "quit" {
			break
		}

		input := scanner.Text()
		value, err := strconv.ParseFloat(input, 64)

		if err != nil {
			fmt.Println("Input Error!")
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

func ReadWall(scanner *bufio.Scanner) (square, string) {
	fmt.Println()
	instruct("wall")
	height, command := InputValue(scanner, "height", "wall")
	if command != "" {
		return square{height: 0, width: 0}, command
	}
	width, command := InputValue(scanner, "width", "wall")
	if command != "" {
		return square{height: 0, width: 0}, command
	}
	isDoubled, command := InputBoolean(scanner, "Paint the wall on both sides?")
	if command != "" {
		return square{height: 0, width: 0}, command
	}

	return square{height: height, width: width, isDoubled: isDoubled}, command
}

func instruct(item string) {
	fmt.Printf("Describe the %s in UNTS (for ex \"1.33\" or \"3\"). Type \"cancel\" to cancel this wall. Type \"quit\" to move to the next step:\n",
		item)
}
