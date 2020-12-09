package common

func FindSum2(target int, numbers []int) (int, int) {
	numMap := make(map[int]bool, len(numbers))
	for _, num := range numbers {
		want := target - num
		if _, ok := numMap[want]; ok {
			return num, want
		}

		numMap[num] = true
	}

	return 0, 0
}

func FindSumContN(target int, numbers []int) []int {
	set := make([]int, 0, len(numbers))
	sum := 0

	index := 0
	indexPtr := 0
	indexMax := len(numbers)

	for {
		if index >= indexMax {
			return nil
		}

		num := numbers[index]

		if num > target {
			indexPtr = index + 1
			index = indexPtr

			sum = 0
			set = make([]int, 0, len(numbers))

			continue
		}

		sum += num

		if sum > target {
			indexPtr++
			index = indexPtr

			sum = 0
			set = make([]int, 0, len(numbers))

			continue
		}

		set = append(set, num)

		if sum == target {
			return set
		}

		index++
	}
}

func FindMinMax(list []int) (min, max int) {
	min = int(^uint(0) >> 1)

	for _, num := range list {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return
}
