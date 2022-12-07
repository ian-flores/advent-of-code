package main

import (
	"bufio"
	"fmt"
	"os"

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

	input_string := fileLines[0]
	input_rune := []rune(input_string)

	pattern_found, character := pattern_finder(14, input_rune)
	fmt.Println(pattern_found, character)
}

func pattern_finder(message_size int, input []rune) (bool, int) {
	pattern_found := false
	idx := 0
	for i := 0; i < len(input); i++ {
		if !pattern_found {
			msg_1 := input[i : i+message_size]
			m := make(map[rune]bool)
			for j := 0; j < len(msg_1); j++ {
				m[msg_1[j]] = true
			}
			if len(m) == message_size {
				idx = i + message_size
				pattern_found = true
			}
		}
	}
	return pattern_found, idx
}
