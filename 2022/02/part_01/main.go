package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ian-flores/advent-of-code/2022/utils"
)

func main() {
	file, err := os.Open("../data.txt")
	defer file.Close()

	utils.ErrorHandler(err)

	scanner := bufio.NewScanner(file)

	total_score := 0

	for scanner.Scan() {
		line := scanner.Text()
		split_line := strings.Split(line, " ")

		play_01 := play_converter(split_line[0])
		play_02 := play_converter(split_line[1])

		game_result := play_winner(play_01, play_02)

		result_score := result_scorer(game_result)
		play_score := play_scorer(play_02)

		total_score += result_score + play_score
	}
	fmt.Println("total score:", total_score)
}

func play_converter(play string) string {
	if play == "A" || play == "X" {
		return "rock"
	} else if play == "B" || play == "Y" {
		return "paper"
	}
	return "scissors"
}

func play_winner(play_01 string, play_02 string) string {
	if play_01 == "rock" && play_02 == "scissors" {
		return "lose"
	} else if play_01 == "rock" && play_02 == "paper" {
		return "win"
	} else if play_01 == "paper" && play_02 == "rock" {
		return "lose"
	} else if play_01 == "paper" && play_02 == "scissors" {
		return "win"
	} else if play_01 == "scissors" && play_02 == "paper" {
		return "lose"
	} else if play_01 == "scissors" && play_02 == "rock" {
		return "win"
	}
	return "tie"
}

func play_scorer(play string) int {
	if play == "rock" {
		return 1
	} else if play == "paper" {
		return 2
	}
	return 3
}

func result_scorer(result string) int {
	if result == "win" {
		return 6
	} else if result == "lose" {
		return 0
	}
	return 3
}
