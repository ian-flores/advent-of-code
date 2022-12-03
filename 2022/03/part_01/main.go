package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/ian-flores/advent-of-code/2022/utils"
)

func main() {
	file, err := os.Open("../data.txt")
	defer file.Close()

	utils.ErrorHandler(err)

	scanner := bufio.NewScanner(file)

	total_count := 0

	for scanner.Scan() {
		line := scanner.Text()
		//line := "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"
		split_string := strings.Split(line, "")
		mid_point := len(line) / 2
		string_1 := split_string[:mid_point]
		string_2 := split_string[mid_point:]

		match := ""
		for _, char_01 := range string_1 {
			for _, char_02 := range string_2 {
				if char_01 == char_02 && char_01 != match {
					match = char_01
					char := []rune(char_01)[0]
					int_value := rune_converter(char)
					total_count += int_value
				}
			}
		}
	}
	fmt.Println("total count:", total_count)
}

func rune_converter(rune rune) int {
	if unicode.IsLower(rune) {
		return int(rune) - 96
	}
	return int(rune) - 38
}
