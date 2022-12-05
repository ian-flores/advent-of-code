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

	var fileLines []string

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	var stacks [9]Stack

	stacks[0] = Stack{"N", "R", "J", "T", "Z", "B", "D", "F"}
	stacks[1] = Stack{"H", "J", "N", "S", "R"}
	stacks[2] = Stack{"Q", "F", "Z", "G", "J", "N", "R", "C"}
	stacks[3] = Stack{"Q", "T", "R", "G", "N", "V", "F"}
	stacks[4] = Stack{"F", "Q", "T", "L"}
	stacks[5] = Stack{"N", "G", "R", "B", "Z", "W", "C", "Q"}
	stacks[6] = Stack{"M", "H", "N", "S", "L", "C", "F"}
	stacks[7] = Stack{"J", "T", "M", "Q", "N", "D"}
	stacks[8] = Stack{"S", "G", "P"}

	for i := 0; i < len(stacks); i++ {
		for j, h := 0, len(stacks[i])-1; j < h; j, h = j+1, h-1 {
			stacks[i][j], stacks[i][h] = stacks[i][h], stacks[i][j]
		}
	}

	for line := 10; line < len(fileLines); line++ {
		message := fileLines[line]
		boxes, giving_stack, receiving_stack := message_parser(message)

		for i := 0; i < boxes; i++ {
			stacks[receiving_stack-1].Push(stacks[giving_stack-1].Pop())
		}
	}

	for i := 0; i < 9; i++ {
		fmt.Println(stacks[i])
	}
}

func message_parser(message string) (int, int, int) {
	split_text := strings.Split(message, " ")
	return string_converter(split_text[1]), string_converter(split_text[3]), string_converter(split_text[5])
}

func string_converter(str string) int {
	integer, _ := strconv.Atoi(str)
	return integer
}

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element
	}
}
