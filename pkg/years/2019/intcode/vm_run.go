package intcode

import (
	"context"
	"fmt"
	"time"
)

func (vm *VM) tick() (status Status) {
	instruction := vm.memory[vm.pc]
	op := instruction % 100
	modes := instruction / 100

	ins := Instruction{
		Op: op,
		Modes: [3]int{
			modes % 10,
			(modes / 10) % 10,
			(modes / 100) % 10,
		},
	}

	return vm.execute(ins)
}

func (vm *VM) Run() {
	ctx, cc := context.WithTimeout(context.Background(), time.Second)
	defer cc()

	defer func() {
		r := recover()
		if r == nil {
			return
		}

		fmt.Println("PC", vm.pc)
		fmt.Println("ROM", vm.rom)
		fmt.Println("ROM_L", len(vm.rom))
		fmt.Println("MEM", vm.memory[:len(vm.rom)+50])
		fmt.Println("IO  IN_L", len(vm.io.in))
		fmt.Println("IO OUT_L", len(vm.io.out))

		panic(r)
	}()

	for {
		select {
		case <-ctx.Done():
			panic("timeout")
		default:
			status := vm.tick()

			if status == StatusHalted {
				vm.halted <- true
				return
			}
		}
	}
}

func (vm *VM) RunUntilInput() {
	for {
		status := vm.tick()

		if status == StatusWaitingForInput {
			return
		}

		if status == StatusHalted {
			vm.halted <- true
			return
		}
	}
}
