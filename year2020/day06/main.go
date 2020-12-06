package main

import (
	"bufio"
	"bytes"
	"fmt"

	"../../utils"
)

func main() {
	data, err := utils.GetInput(2020, 06)
	utils.CheckErr(err)

	scanner := bufio.NewScanner(bytes.NewBuffer(data))
	scanner.Split(bufio.ScanLines)

	ans := make(map[rune]int)
	total := 0
	totalUnique := 0
	peopleCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			total += len(ans)
			for _, val := range ans {
				if val == peopleCount {
					totalUnique++
				}
			}

			ans = make(map[rune]int)
			peopleCount = 0

			continue
		}
		peopleCount++

		for _, letter := range line {
			val := ans[letter]
			ans[letter] = val + 1
		}
	}

	total += len(ans)
	for _, val := range ans {
		if val == peopleCount {
			totalUnique++
		}
	}

	fmt.Println("total", total)
	fmt.Println("total", totalUnique)
}
