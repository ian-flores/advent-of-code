package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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

	//numbers_slice := inputLines[0]

	playerBoards := buildBoards(inputLines[1:])
	fmt.Printf("%v", playerBoards)

}

func buildBoards(inputs []string) [][][]string {
	boards := make([][][]string, len(inputs)/6)
	board_counter := 0
	for _, line := range inputs {
		row_counter := 0
		if line == "" {
			board_counter++
			continue
		} else {
			line_split := strings.Split(line, " ")
			for column_idx, row_value := range line_split {
				boards[row_counter][column_idx][board_counter] = row_value
			}
			row_counter++
		}

	}

	return boards
}

//func checkBoard(target_number string, player_board []string, player_board_state []int) (bool, []string, error) {
//	return true, player_board, nil
//}
