package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	dialStart  = 0
	dialEnd    = 99
	initialPos = 50
)

func main() {
	file, _ := os.Open("data/day-01_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	currentPosition := initialPos
	counterPart1, counterPart2 := 0, 0

	for scanner.Scan() {
		text := scanner.Text()

		var zeroCounter int
		currentPosition, zeroCounter = moveDial(currentPosition, text[0], text[1:])

		if currentPosition == 0 {
			counterPart1++
		}

		counterPart2 += zeroCounter
	}

	fmt.Println("Part 1 Counter: ", counterPart1)
	fmt.Println("Part 2 Counter: ", counterPart2)

}

func moveDial(currentPosition int, direction byte, clicks string) (int, int) {
	zeroCounter := 0

	clicksInt, _ := strconv.Atoi(clicks)
	for range clicksInt {
		switch direction {
		case 'L':
			currentPosition--
			if currentPosition < dialStart {
				currentPosition = dialEnd
			}
		case 'R':
			currentPosition++
			if currentPosition > dialEnd {
				currentPosition = dialStart
			}
		}

		if currentPosition == 0 {
			zeroCounter++
		}
	}

	return currentPosition, zeroCounter
}
