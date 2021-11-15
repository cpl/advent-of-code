package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"../../utils"
)

type Instruction struct {
	OP    string
	Value int
}

type VM struct {
	ACC int
	PC  int

	Instructions []*Instruction
}

func (vm *VM) Run(iter int, print bool) bool {
	vm.PC = 0
	vm.ACC = 0

	maxPC := len(vm.Instructions)
	hitCount := make([]int, maxPC)
	nopjmpCount := 0

	for {
		if vm.PC == maxPC {
			return true
		}

		ins := vm.Instructions[vm.PC]

		hc := hitCount[vm.PC]
		if hc == 1 {
			return false
		}
		hitCount[vm.PC] = hc + 1

		if print {
			fmt.Printf("%4d %6s %6d\n", vm.PC, ins.OP, ins.Value)
		}

		switch ins.OP {
		case "nop":
			nopjmpCount++
			if nopjmpCount == iter {
				vm.PC += ins.Value
				fmt.Printf("%4d %6s %6d - replace %4d\n", vm.PC, ins.OP, ins.Value, iter)
				continue
			}
		case "jmp":
			nopjmpCount++
			if nopjmpCount == iter {
				fmt.Printf("%4d %6s %6d - replace %4d\n", vm.PC, ins.OP, ins.Value, iter)
			} else {
				vm.PC += ins.Value
				continue
			}
		case "acc":
			vm.ACC += ins.Value
		default:
			fmt.Println("panic, unknown op")
			return false
		}

		vm.PC++
	}
}

func main() {
	data, err := utils.GetInput(2020, 8)
	utils.CheckErr(err)

	scanner := bufio.NewScanner(bytes.NewBuffer(data))
	scanner.Split(bufio.ScanLines)

	vm := VM{
		Instructions: make([]*Instruction, 0, 1000),
	}

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		value, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		vm.Instructions = append(vm.Instructions, &Instruction{
			OP:    split[0],
			Value: value,
		})
	}

	ok := vm.Run(0, true)
	fmt.Println("ACC", vm.ACC, ok)

	iter := 1
	for {
		if vm.Run(iter, false) {
			break
		}
		iter++
	}
	fmt.Println("replaced #", iter, "occurrence of JMP or NOP, final ACC", vm.ACC)
}
