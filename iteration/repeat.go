// Practice with for loops and iteration
package iteration

// Repeats a string, str, a number of times, count.
func Repeat(str string, count int) string {
	var repeatedStr string
	for i := 0; i < count; i++ {
		repeatedStr += str
	}
	return repeatedStr
}
