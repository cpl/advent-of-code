package y2025_test

import (
	"bufio"
	"bytes"
	"slices"
	"strconv"
	"testing"

	"github.com/cpl/advent-of-code/pkg/aoc"
	aocparse "github.com/cpl/advent-of-code/pkg/aoc-parse"
	"go.sdls.io/pkg"
)

func TestSolve2025Day06(t *testing.T) {
	aoc.SolveExample(t, "example", 1, "4277556", "123 328  51 64 \n 45 64  387 23 \n  6 98  215 314\n*   +   *   +  ")
	aoc.SolveExample(t, "example", 2, "3263827", "123 328  51 64 \n 45 64  387 23 \n  6 98  215 314\n*   +   *   +  ")

	aoc.Solve(t, "part 1", aocparse.ColumnsFields(), func(columns [][]string) string {
		total := pkg.UInt128{}

		for _, column := range columns {
			subtotal := pkg.UInt128{}

			symbol := column[len(column)-1]
			switch symbol {
			case "+":
				for _, cell := range column[:len(column)-1] {
					num, _ := strconv.ParseUint(cell, 10, 64)
					subtotal = subtotal.Add64(num)
				}
			case "*":
				subtotal = subtotal.Add64(1)
				for _, cell := range column[:len(column)-1] {
					num, _ := strconv.ParseUint(cell, 10, 64)
					subtotal = subtotal.Mul64(num)
				}

				if subtotal.Compare64(1) == 0 {
					subtotal = pkg.UInt128{}
				}
			}

			total = total.Add(subtotal)
		}

		return total.String()
	})

	cephalopodMathParser := func(scanner *bufio.Scanner) [][]byte {
		scanner.Split(bufio.ScanLines)

		if !scanner.Scan() {
			panic("empty scanner?")
		}
		line := scanner.Bytes()
		out := make([][]byte, len(line))
		for idx, ch := range line {
			out[idx] = append(out[idx], ch)
		}

		for scanner.Scan() {
			for idx, ch := range scanner.Bytes() {
				out[idx] = append(out[idx], ch)
			}
		}

		return out
	}

	aoc.Solve(t, "part 2", cephalopodMathParser, func(chars [][]byte) string {
		numbers := make([]uint64, 0, 64)
		buf := bytes.NewBuffer(nil)
		symbol := byte(0)

		total := pkg.UInt128{}

		chars = slices.Insert(chars, 0, []byte{' '})
		for _, line := range slices.Backward(chars) {
			buf.Reset()
			allSpace := true

			for _, ch := range line {
				if ch == ' ' {
					if buf.Len() > 0 {
						num, _ := strconv.ParseUint(buf.String(), 10, 64)
						numbers = append(numbers, num)
						buf.Reset()
					}

					continue
				}

				allSpace = false

				isDigit := ch >= '0' && ch <= '9'
				if isDigit {
					buf.WriteByte(ch)
					continue
				}

				isSymbol := ch == '+' || ch == '*'
				if isSymbol {
					symbol = ch

					if buf.Len() > 0 {
						num, _ := strconv.ParseUint(buf.String(), 10, 64)
						numbers = append(numbers, num)
						buf.Reset()
					}

					continue
				}
			}

			if !allSpace {
				continue
			}

			if buf.Len() > 0 {
				num, _ := strconv.ParseUint(buf.String(), 10, 64)
				numbers = append(numbers, num)
				buf.Reset()
			}

			switch symbol {
			case '+':
				subtotal := pkg.UInt128{}
				for _, num := range numbers {
					subtotal = subtotal.Add64(num)
				}

				numbers = numbers[:0]
				total = total.Add(subtotal)
			case '*':
				subtotal := pkg.UInt128From64(1)
				for _, num := range numbers {
					subtotal = subtotal.Mul64(num)
				}

				if subtotal.Compare64(1) == 0 {
					subtotal = pkg.UInt128{}
				}

				numbers = numbers[:0]
				total = total.Add(subtotal)
			default:
				panic("unreachable")
			}
		}

		return total.String()
	})
}
