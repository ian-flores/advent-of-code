package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	window := flag.Int("window", 1, "Window Size to Scan")
	inputFile := flag.String("inputFile", "./input.txt", "Path to the Input File")

	flag.Parse()

	file, err := os.Open(*inputFile)

	if err != nil {
		log.Fatal("Couldn't open the input file:", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var inputLines []int

	for scanner.Scan() {
		value := scanner.Text()
		intValue, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("Couldn't convert one of the lines to an integer value:", err)
		}

		inputLines = append(inputLines, intValue)
	}

	file.Close()

	counter := 0

	// From: https://github.com/orfeasa/advent-of-code-2021/blob/91a8f9e6034f39fc58a60ce34866b201a6975a32/day_01/go/main.go#L27
	for i := 0; i < len(inputLines)-*window; i++ {
		if inputLines[i+*window] > inputLines[i] {
			counter++
		}
	}

	// Original Try
	// for idx, line := range inputLines {
	// 	if idx == 0 {
	// 		previous_value = line
	// 		continue
	// 	}
	// 	if line > previous_value {
	// 		counter++
	// 		previous_value = line
	// 	}
	// }

	fmt.Println("The total number of lines sufficing your window slider is:", counter)
}
