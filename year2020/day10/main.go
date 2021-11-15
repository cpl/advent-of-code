package main

import (
	"bufio"
	"bytes"
	"fmt"
	"sort"
	"strconv"

	"../../utils"
)

const lastSize = 25

func main() {
	data, err := utils.GetInput(2020, 10)
	utils.CheckErr(err)

	scanner := bufio.NewScanner(bytes.NewBuffer(data))
	scanner.Split(bufio.ScanLines)

	jolts := make([]int, 0, 1000)

	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		jolts = append(jolts, num)
	}

	sort.IntSlice(jolts).Sort()
	jolts = append(jolts, jolts[len(jolts)-1]+3)

	currentJolt := 0
	joltDiff := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
	}
	joltDiffList := make([]int, len(jolts))

	for idx, jolt := range jolts {
		if jolt-currentJolt > 3 {
			fmt.Printf("jolt %d at index %d creates gap\n", jolt, idx)
			break
		}

		joltDiffList[idx] = jolt - currentJolt
		joltDiff[jolt-currentJolt] += 1
		currentJolt = jolt
	}

	fmt.Println(jolts[0], jolts[len(jolts)-1])
	fmt.Println(joltDiff)
	fmt.Println(joltDiff[1] * joltDiff[3])
	fmt.Println(joltDiffList)

	count1 := 0
	countSeq := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
	}
	for _, diff := range joltDiffList {
		if diff == 3 {
			if count1 > 0 {
				fmt.Println()
			}

			countSeq[count1] += 1
			count1 = 0
			continue
		}
		count1++
		fmt.Print(diff)
		fmt.Print(" ")
	}
	fmt.Println(countSeq)
}
