package intcode

func (vm *VM) IO() (in, out chan int) {
	return vm.io.in, vm.io.out
}

func (vm *VM) IOWrite(values ...int) {
	for _, v := range values {
		vm.io.in <- v
	}
}

func (vm *VM) IORead() []int {
	l := len(vm.io.out)
	if l == 0 {
		return nil
	}

	values := make([]int, l)

	for i := 0; i < l; i++ {
		values[i] = <-vm.io.out
	}

	return values
}

func (vm *VM) IOReadBlocking() int {
	return <-vm.io.out
}
