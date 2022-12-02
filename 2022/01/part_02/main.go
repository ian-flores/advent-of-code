package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"container/heap"

	"github.com/ian-flores/advent-of-code/2022/utils"
)

func main() {
	file, err := os.Open("../data.txt")
	defer file.Close()

	utils.ErrorHandler(err)

	scanner := bufio.NewScanner(file)

	deer_counter, total_cals := 0, 0

	cals_map := make(map[string]int)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			deer_counter += 1
		} else {
			calories, err := strconv.Atoi(text)
			utils.ErrorHandler(err)

			deer_id := strconv.Itoa(deer_counter)

			cals_map[deer_id] = cals_map[deer_id] + calories
		}
	}

	utils.ErrorHandler(scanner.Err())

	h := utils.GetHeap(cals_map)

	for i := 0; i <= 2; i++ {
		value := heap.Pop(h).(utils.KV).Value
		total_cals += value
	}

	fmt.Println("total cals:", total_cals)
}
