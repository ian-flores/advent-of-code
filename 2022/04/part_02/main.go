package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ian-flores/advent-of-code/2022/utils"
)

func main() {
	file, err := os.Open("../data.txt")
	defer file.Close()

	utils.ErrorHandler(err)

	scanner := bufio.NewScanner(file)

	total_count := 0

	for scanner.Scan() {
		text := scanner.Text()
		ranges := strings.Split(text, ",")
		range_01 := strings.Split(ranges[0], "-")
		range_02 := strings.Split(ranges[1], "-")
		if range_comparer(range_01, range_02) {
			total_count += 1
		}
	}
	fmt.Println("total count:", total_count)
}

func range_comparer(range_01 []string, range_02 []string) bool {
	if range_checker(range_01, range_02[0]) || range_checker(range_01, range_02[1]) {
		return true
	} else if range_checker(range_02, range_01[0]) || range_checker(range_02, range_01[1]) {
		return true
	}
	return false
}

func range_checker(ranges []string, number_input string) bool {
	number, err := strconv.Atoi(number_input)
	utils.ErrorHandler(err)
	min, err := strconv.Atoi(ranges[0])
	utils.ErrorHandler(err)
	max, err := strconv.Atoi(ranges[1])
	utils.ErrorHandler(err)

	if number >= min && number <= max {
		return true
	}
	return false
}
