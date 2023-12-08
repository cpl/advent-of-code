package aoc_math

func LCM(numbers ...int64) int64 {
	var lcm int64 = 1

	for _, n := range numbers {
		lcm = lcm * n / GCD(lcm, n)
	}

	return lcm
}

func GCD(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
