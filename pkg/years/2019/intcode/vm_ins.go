package intcode

import (
	"fmt"
)

type Instruction struct {
	Op    int
	Modes [3]int
}

type Status byte

const (
	StatusUndefined Status = iota
	StatusHalted
	StatusWaitingForInput
	StatusRunning
)

func (vm *VM) execute(ins Instruction) (status Status) {
	switch ins.Op {
	case 1:
		rA := vm.load(vm.pc+1, ins.Modes[0])
		RB := vm.load(vm.pc+2, ins.Modes[1])
		vm.store(vm.pc+3, rA+RB, ins.Modes[2])

		vm.pc += 4
	case 2:
		rA := vm.load(vm.pc+1, ins.Modes[0])
		rB := vm.load(vm.pc+2, ins.Modes[1])
		vm.store(vm.pc+3, rA*rB, ins.Modes[2])

		vm.pc += 4
	case 3:
		select {
		case rA := <-vm.io.in:
			vm.store(vm.pc+1, rA, ins.Modes[0])
		default:
			return StatusWaitingForInput
		}

		vm.pc += 2
	case 4:
		rA := vm.load(vm.pc+1, ins.Modes[0])
		vm.io.out <- rA

		vm.pc += 2
	case 5:
		rA := vm.load(vm.pc+1, ins.Modes[0])
		if rA != 0 {
			vm.pc = vm.load(vm.pc+2, ins.Modes[1])
			return StatusRunning
		}

		vm.pc += 3
	case 6:
		rA := vm.load(vm.pc+1, ins.Modes[0])
		if rA == 0 {
			vm.pc = vm.load(vm.pc+2, ins.Modes[1])
			return StatusRunning
		}

		vm.pc += 3
	case 7:
		rA := vm.load(vm.pc+1, ins.Modes[0])
		rB := vm.load(vm.pc+2, ins.Modes[1])
		if rA < rB {
			vm.store(vm.pc+3, 1, ins.Modes[2])
		} else {
			vm.store(vm.pc+3, 0, ins.Modes[2])
		}

		vm.pc += 4
	case 8:
		rA := vm.load(vm.pc+1, ins.Modes[0])
		rB := vm.load(vm.pc+2, ins.Modes[1])
		if rA == rB {
			vm.store(vm.pc+3, 1, ins.Modes[2])
		} else {
			vm.store(vm.pc+3, 0, ins.Modes[2])
		}

		vm.pc += 4
	case 9:
		rA := vm.load(vm.pc+1, ins.Modes[0])
		vm.rb += rA

		vm.pc += 2
	case 99:
		return StatusHalted
	default:
		panic(fmt.Sprintf("unknown opcode %d (%+v)", ins.Op, ins))
	}

	return StatusRunning
}
