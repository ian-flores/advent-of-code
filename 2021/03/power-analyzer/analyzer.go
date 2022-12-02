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
	var columns_total [12]int
	for _, line := range inputLines {
		line_slice := strings.Split(line, "")
		for i := 0; i < len(line_slice); i++ {
			column_value, _ := strconv.Atoi(line_slice[i])
			columns_total[i] += column_value
		}
	}

	var gamma_number []string
	var epsilon_number []string

	for i := 0; i < len(columns_total); i++ {
		if columns_total[i] > len(inputLines)/2 {
			gamma_number = append(gamma_number, "1")
			epsilon_number = append(epsilon_number, "0")
		} else {
			gamma_number = append(gamma_number, "0")
			epsilon_number = append(epsilon_number, "1")
		}
	}

	gamma_number_string := strings.Join(gamma_number, "")
	epsilon_number_string := strings.Join(epsilon_number, "")

	decimal_gamma_number, _ := strconv.ParseInt(gamma_number_string, 2, 64)
	decimal_epsilon_number, _ := strconv.ParseInt(epsilon_number_string, 2, 64)

	fmt.Println("Gamma Rate:", decimal_gamma_number)
	fmt.Println("Epsilon Rate:", decimal_epsilon_number)
	fmt.Println("Total Power Rate:", decimal_gamma_number*decimal_epsilon_number)
}
