package main

import (
	"fmt"
	"io/ioutil"
	"unicode"
)

type stack struct {
	data [50000]rune
	SP   int
}

func (s *stack) PUSH(value rune) {
	s.data[s.SP] = value
	s.SP++
}

func (s *stack) POP() rune {
	s.SP--
	return s.data[s.SP]
}

func (s *stack) PEAK() rune {
	if s.SP == 0 {
		return 0
	}
	return s.data[s.SP-1]
}

var alpha = []rune{
	'a', 'b', 'c', 'd', 'e',
	'f', 'g', 'h', 'i', 'j',
	'k', 'l', 'm', 'n', 'o',
	'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z'}

func main() {
	var s stack
	s.SP = 0

	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	minSP := 50000

	for _, ignore := range alpha {
		s.SP = 0
		fmt.Printf("Ignoring: %c\n", ignore)

		for _, c := range data {
			r := rune(c)

			if unicode.ToLower(r) == ignore {
				continue
			}

			if unicode.IsLower(r) && unicode.ToUpper(r) == s.PEAK() {
				s.POP()
			} else if unicode.IsUpper(r) && unicode.ToLower(r) == s.PEAK() {
				s.POP()
			} else {
				s.PUSH(r)
			}
		}

		fmt.Printf("SP: %d\n", s.SP)
		if s.SP < minSP {
			minSP = s.SP
		}
	}

	fmt.Println("min SP: ", minSP)
}
