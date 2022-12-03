package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	"github.com/ian-flores/advent-of-code/2022/utils"
)

func main() {
	file, err := os.Open("../data.txt")
	defer file.Close()

	utils.ErrorHandler(err)

	scanner := bufio.NewScanner(file)

	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	total_count := 0

	for i := 0; i < len(fileLines); i = i + 3 {
		string_1 := fileLines[i]
		string_2 := fileLines[i+1]
		string_3 := fileLines[i+2]

		var match rune
		for _, char_01 := range string_1 {
			for _, char_02 := range string_2 {
				for _, char_03 := range string_3 {
					if char_01 == char_02 && char_01 == char_03 && char_01 != match {
						match = char_01
						int_value := rune_converter(char_01)
						total_count += int_value
					}
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
