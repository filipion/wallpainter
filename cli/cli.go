package cli

import (
	"bufio"
	"fmt"
	"strconv"
)

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
