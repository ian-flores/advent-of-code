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
		game_result := result_converter(split_line[1])
		play_02 := play_generator(play_01, game_result)

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

func result_converter(result string) string {
	if result == "X" {
		return "lose"
	} else if result == "Y" {
		return "tie"
	}
	return "win"
}

func play_generator(play_01 string, game_result string) string {
	if play_01 == "rock" && game_result == "win" {
		return "paper"
	} else if play_01 == "rock" && game_result == "lose" {
		return "scissors"
	} else if play_01 == "rock" && game_result == "tie" {
		return "rock"
	} else if play_01 == "paper" && game_result == "win" {
		return "scissors"
	} else if play_01 == "paper" && game_result == "lose" {
		return "rock"
	} else if play_01 == "paper" && game_result == "tie" {
		return "paper"
	} else if play_01 == "scissors" && game_result == "win" {
		return "rock"
	} else if play_01 == "scissors" && game_result == "lose" {
		return "paper"
	}
	return "scissors"
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
