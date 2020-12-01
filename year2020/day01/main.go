package main

import (
	"../../utils"
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	data, err := utils.GetInput(2020, 01)
	checkErr(err)

	scanner := bufio.NewScanner(bytes.NewBuffer(data))
	scanner.Split(bufio.ScanLines)

	target := 2020
	numbers := make([]int, 0)

	for scanner.Scan() {
		fmt.Print(".")
		line := strings.TrimSpace(scanner.Text())

		num, _ := strconv.Atoi(line)
		numbers = append(numbers, num)
	}
	fmt.Println()

	part1A, part1B := findSum2(target, numbers)
	log.Println("result part1:", part1A, part1B, part1A*part1B)

	part2A, part2B, part2C := findSum3(target, numbers)
	log.Println("result part2:", part2A, part2B, part2C, part2A*part2B*part2C)
}

func findSum3(target int, numbers []int) (int, int, int) {
	for idx, num := range numbers {
		a, b := findSum2(target-num, numbers[idx:])
		if a != 0 && b != 0 {
			return num, a, b
		}
	}

	return 0, 0, 0
}

func findSum2(target int, numbers []int) (int, int) {
	numMap := make(map[int]bool, len(numbers))
	for _, num := range numbers {
		want := target-num
		if _, ok := numMap[want]; ok {
			return num, want
		}

		numMap[num] = true
	}

	return 0, 0
}
