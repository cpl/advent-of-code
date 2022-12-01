package intcode

type VM struct {
	pc int
	rb int

	rom    []int
	memory []int

	io struct {
		in  chan int
		out chan int
	}

	halted chan bool
}

const (
	configMaxMemory = 1 << 14
	configMaxIOIn   = 100
	configMaxIOOut  = 10000
)

func (vm *VM) Bootstrap(program []int) {
	vm.io.in = make(chan int, configMaxIOIn)
	vm.io.out = make(chan int, configMaxIOOut)
	vm.halted = make(chan bool, 1)

	vm.rom = program
	vm.memory = make([]int, configMaxMemory)

	copy(vm.memory, program)
}

func (vm *VM) Reset() {
	close(vm.io.in)
	vm.io.in = make(chan int, configMaxIOIn)
	close(vm.io.out)
	vm.io.out = make(chan int, configMaxIOOut)
	close(vm.halted)
	vm.halted = make(chan bool, 1)

	copy(vm.memory, vm.rom)
	vm.Memclear(len(vm.rom), len(vm.memory)-len(vm.rom))

	vm.pc = 0
	vm.rb = 0
}

func (vm *VM) Memset(idx, value int) {
	vm.memory[idx] = value
}

func (vm *VM) Memget(idx int) int {
	return vm.memory[idx]
}

func (vm *VM) Memclear(from, size int) {
	if from+size > len(vm.memory) {
		size = len(vm.memory) - from
	}

	for i := from; i < from+size; i++ {
		vm.memory[i] = 0
	}
}

func (vm *VM) Halted() chan bool {
	return vm.halted
}
