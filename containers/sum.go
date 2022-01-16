package containers

func Sum(vals [4]int) int {
	var sum int
	for _, number := range vals {
		sum += number
	}
	return sum
}
