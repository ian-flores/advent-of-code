package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../data.txt")

	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	max_cals := 0
	deer_cals := 0
	deer_counter := 0

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			deer_cals = 0
			deer_counter += 1
		} else {
			cals, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
			deer_cals = deer_cals + cals
			if deer_cals > max_cals {
				max_cals = deer_cals
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("max cals:", max_cals)
	fmt.Println("deer counter:", deer_counter)
}
