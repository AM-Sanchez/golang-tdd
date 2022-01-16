package containers

func Sum(vals []int) int {
	var sum int
	for _, number := range vals {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {

	for _, intSlice := range numbersToSum {
		sum := 0
		for _, val := range intSlice {
			sum += val
		}
		sums = append(sums, sum)
	}
	return sums
}
