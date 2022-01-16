package containers

func Sum(vals []int) int {
	var sum int
	for _, number := range vals {
		sum += number
	}
	return sum
}
