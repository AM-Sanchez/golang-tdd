package containers

func Sum(vals []int) int {
	var sum int
	for _, number := range vals {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numberSlice := range numbersToSum {
		sums = append(sums, Sum(numberSlice))
	}

	return sums
}

func SumAllTails(numbersToSumTails ...[]int) (tailSums []int) {
	for _, numberSlice := range numbersToSumTails {
		if len(numberSlice) == 0 {
			// Handle an empty slice
			tailSums = append(tailSums, 0)
		} else {
			// Create the tail values to sum by excluding the first element of the slice
			tailValues := numberSlice[1:]
			tailSums = append(tailSums, Sum(tailValues))
		}
	}
	return tailSums
}
