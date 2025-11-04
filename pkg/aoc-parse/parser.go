package aoc_parse

import "bufio"

type Parser[T any] = func(scanner *bufio.Scanner) T
