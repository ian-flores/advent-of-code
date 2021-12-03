package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile := flag.String("inputFile", "./input.txt", "Path to the Input File")

	flag.Parse()

	file, err := os.Open(*inputFile)

	if err != nil {
		log.Fatal("Couldn't open the input file:", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var inputLines []string

	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}

	file.Close()

	horizontal_position := 0
	depth := 0
	aim := 0

	for _, line := range inputLines {
		current_movement := strings.Split(line, " ")
		movementChange, err := strconv.Atoi(current_movement[1])

		if err != nil {
			log.Fatal("Couldn't convert one of the lines to an integer value:", err)
		}
		if current_movement[0] == "forward" {
			horizontal_position = horizontal_position + movementChange
			depth = depth + aim*movementChange
		}
		if current_movement[0] == "up" {
			aim = aim - movementChange
		}
		if current_movement[0] == "down" {
			aim = aim + movementChange
		}
	}

	fmt.Println("Final position:", horizontal_position, depth)
	fmt.Println("Total Multiplication", horizontal_position*depth)
}
