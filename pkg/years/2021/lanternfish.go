package y2021

import (
	"strconv"
	"strings"
)

type LanternFishBank struct {
	fish [9]int
}

func (bank *LanternFishBank) Age() (spawn int) {
	spawn = bank.fish[0]

	for idx := 0; idx < 8; idx++ {
		bank.fish[idx] = bank.fish[idx+1]
	}

	bank.fish[6] += spawn
	bank.fish[8] = spawn

	return
}

func (bank *LanternFishBank) Total() int {
	total := 0
	for _, f := range bank.fish {
		total += f
	}
	return total
}

func (bank *LanternFishBank) AgeDays(days int) {
	for day := 0; day < days; day++ {
		_ = bank.Age()
	}
}

func ParseLanternFishBank(input string) *LanternFishBank {
	input = strings.TrimSpace(input)
	bank := &LanternFishBank{}
	for _, day := range strings.Split(input, ",") {
		dayInt, _ := strconv.Atoi(day)
		bank.fish[dayInt]++
	}

	return bank
}
