package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"

	"../../common"
	"../../utils"
)

const lastSize = 25

func main() {
	data, err := utils.GetInput(2020, 9)
	utils.CheckErr(err)

	scanner := bufio.NewScanner(bytes.NewBuffer(data))
	scanner.Split(bufio.ScanLines)

	last := make([]int, lastSize)
	allN := make([]int, 0, 1000)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()

		num, _ := strconv.Atoi(line)

		if count >= lastSize {
			a, b := common.FindSum2(num, last)
			if a == 0 && b == 0 {
				fmt.Println(num, last)
				sumnums := common.FindSumContN(num, allN)
				min, max := common.FindMinMax(sumnums)

				fmt.Println(sumnums)
				fmt.Println(min, "+", max, "=", min+max)

				return
			}
		}

		allN = append(allN, num)
		last[count%lastSize] = num
		count++
	}
}
