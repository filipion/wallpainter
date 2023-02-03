package square

import (
	"bufio"
	"fmt"

	"example.com/wallpaint/cli"
)

type Square struct {
	height    float64
	width     float64
	isDoubled bool
	unit      string
	windows   []Square
}

func Area(wall Square) float64 {
	area := wall.height * wall.width
	for _, window := range wall.windows {
		area -= window.height * window.width
	}
	if wall.isDoubled {
		area *= 2
	}
	return area
}

func Represent(wall Square) string {
	text := fmt.Sprintf("Wall of height %.2f and width %.2f\n", wall.height, wall.width)
	for _, window := range wall.windows {
		text += fmt.Sprintf("  window of height %.2f and width %.2f\n", window.height, window.width)
	}
	return text
}

func ReadItem(scanner *bufio.Scanner, item string) (Square, string) {
	isDoubled := true
	windowList := []Square{}
	fmt.Println()
	Instruct(item)
	height, exitCode := cli.InputValue(scanner, "height", item, Instruct(item))
	if exitCode != "" {
		return Square{height: 0, width: 0}, exitCode
	}
	width, exitCode := cli.InputValue(scanner, "width", item, Instruct(item))
	if exitCode != "" {
		return Square{height: 0, width: 0}, exitCode
	}
	if item == "wall" {
		isDoubled, exitCode = cli.InputBoolean(scanner, "Paint the wall on both sides?")
		if exitCode != "" {
			return Square{height: 0, width: 0}, exitCode
		}

		hasOpenings, command := cli.InputBoolean(scanner, "Does the wall have doors, windows or other rectangular openings?")
		if command != "" {
			return Square{height: 0, width: 0}, command
		}
		if hasOpenings {
			fmt.Println("Please specify the sizes of the windows/doors. Type \"end\" after you are done with windows/doors.")
			for {
				window, exitCode := ReadItem(scanner, "window/door")
				if exitCode == "quit" {
					return Square{}, exitCode
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

	value := Square{height: height, width: width, isDoubled: isDoubled, windows: windowList}
	if item == "wall" {
		fmt.Println("\nCurrent wall results:", Represent(value))
	}
	return value, exitCode
}

func Instruct(item string) string {
	return fmt.Sprintf("\nDescribe the %s in meters (for ex \"1.33\" or \"3\"). Type \"cancel\" to cancel this %s. Type \"end\" to finish inputing %ss:\n",
		item, item, item)
}
