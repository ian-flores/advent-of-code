package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/ian-flores/advent-of-code/2022/utils"
)

func main() {
	file, err := os.Open("../data.txt")
	defer file.Close()

	utils.ErrorHandler(err)

	scanner := bufio.NewScanner(file)

	max_cals, deer_cals, deer_counter := 0, 0, 0

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			deer_cals = 0
			deer_counter += 1
		} else {
			cals, err := strconv.Atoi(text)
			utils.ErrorHandler(err)

			deer_cals += cals

			if deer_cals > max_cals {
				max_cals = deer_cals
			}
		}
	}

	utils.ErrorHandler(scanner.Err())

	fmt.Println("max cals:", max_cals)
	fmt.Println("deer counter:", deer_counter)
}
