package main

import (
	"bufio"
	"bytes"
	"fmt"

	"../../utils"
)

type pswdEntry struct {
	min      int
	max      int
	letter   string
	password string
}

func (e *pswdEntry) IsValid1() bool {
	l := len(e.password)
	if l < e.min {
		return false
	}

	letterCount := 0
	for _, char := range []rune(e.password) {
		if string(char) == e.letter {
			letterCount++
		}

		if letterCount > e.max {
			return false
		}
	}

	return letterCount >= e.min
}

func (e *pswdEntry) IsValid2() bool {
	r1 := string(e.password[e.min-1])
	r2 := string(e.password[e.max-1])

	if r1 == e.letter {
		if r2 == e.letter {
			return false
		}
		return true
	} else {
		if r2 == e.letter {
			return true
		}
		return false
	}
}

func main() {
	data, err := utils.GetInput(2020, 02)
	utils.CheckErr(err)

	scanner := bufio.NewScanner(bytes.NewBuffer(data))
	scanner.Split(bufio.ScanLines)

	var entry pswdEntry

	valid1 := 0
	valid2 := 0
	for scanner.Scan() {
		line := scanner.Text()

		_, err := fmt.Sscanf(line, "%d-%d %1s: %s", &entry.min, &entry.max, &entry.letter, &entry.password)
		utils.CheckErr(err)

		if entry.IsValid1() {
			valid1++
		}
		if entry.IsValid2() {
			valid2++
		}
	}

	fmt.Println(valid1, valid2)
}
