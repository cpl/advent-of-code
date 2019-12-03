package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func getIntops() []int {
	data, err := ioutil.ReadFile("year2019/day2/input.txt")
	if err != nil {
		panic(err)
	}
	line := string(data)
	opss := strings.Split(line, ",")


	intops := make([]int, len(opss))
	for idx, op := range opss {
		intops[idx], err = strconv.Atoi(strings.TrimSpace(op))
		if err != nil {
			panic(err)
		}
	}

	return intops
}

const (
	OpAdd = 1
	OpMul = 2
	OpStp = 99
)

func runVM(intops []int) []int {
	out := make([]int, len(intops))
	copy(out, intops)

	var pc int
	for pc < len(out) && out[pc] != OpStp{
		switch out[pc] {
		case OpAdd:
			fmt.Println("+", out[pc:pc+4])
			out[out[pc+3]] = out[out[pc+1]] + out[out[pc+2]]
			pc+=4
		case OpMul:
			fmt.Println("*", out[pc:pc+4])
			out[out[pc+3]] = out[out[pc+1]] * out[out[pc+2]]
			pc+=4
		default:
			fmt.Printf("unkown intop %d at PC %d\n", out[pc], pc)
			os.Exit(1)
		}
	}

	//fmt.Println("mem:", intops)
	//fmt.Println("[0_]:", intops[0])
	//fmt.Println("[PC]:", pc)

	return out
}

func main() {
	intops := getIntops()
	fmt.Printf("parsing %d intops\n", len(intops))

	var noun, verb int
	intops[1] = noun
	intops[2] = verb
	out := runVM(intops)

	for out[0] != 19690720 {
		noun++
		if noun == 100 {
			verb += 1
			noun = 0
		}

		intops[1] = noun
		intops[2] = verb
		out = runVM(intops)
	}

	fmt.Println(out)
	fmt.Println(out[0])
	fmt.Println(noun, verb, 100*noun+verb)
}