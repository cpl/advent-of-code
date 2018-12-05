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
		fmt.Println("PEAK with SP 0")
		return 0
	}
	return s.data[s.SP-1]
}

func main() {
	var s stack
	s.SP = 0

	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	for _, c := range data {
		r := rune(c)

		if unicode.IsLower(r) && unicode.ToUpper(r) == s.PEAK() {
			s.POP()
		} else if unicode.IsUpper(r) && unicode.ToLower(r) == s.PEAK() {
			s.POP()
		} else {
			s.PUSH(r)
		}
	}

	fmt.Println("SP: ", s.SP)
}
