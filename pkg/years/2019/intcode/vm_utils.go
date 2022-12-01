package intcode

import "strconv"

func (vm *VM) load(idx, mode int) int {
	switch mode {
	case 0:
		return vm.memory[vm.memory[idx]]
	case 1:
		return vm.memory[idx]
	case 2:
		return vm.memory[vm.rb+vm.memory[idx]]
	}

	panic("invalid mode " + strconv.Itoa(mode))
}

func (vm *VM) store(idx, value, mode int) {
	switch mode {
	case 0:
		vm.memory[vm.memory[idx]] = value
	case 1:
		vm.memory[idx] = value
	case 2:
		vm.memory[vm.rb+vm.memory[idx]] = value
	default:
		panic("invalid mode " + strconv.Itoa(mode))
	}
}
