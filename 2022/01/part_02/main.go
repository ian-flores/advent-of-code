package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"container/heap"

	"github.com/ian-flores/advent-of-code/2022/utils"
)

func main() {
	file, err := os.Open("../data.txt")

	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	deer_counter := 0

	cals_map := make(map[string]int)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			deer_counter += 1
		} else {
			calories, err := strconv.Atoi(text)
			deer_id := strconv.Itoa(deer_counter)
			if err != nil {
				log.Fatal(err)
			}
			cals_map[deer_id] = cals_map[deer_id] + calories
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	h := utils.GetHeap(cals_map)

	total_cals := 0

	for i := 0; i <= 2; i++ {
		value := heap.Pop(h).(utils.KV).Value
		total_cals += value
		fmt.Printf("%d) %#v\n", i, value)
	}

	fmt.Println("total cals:", total_cals)
}
