package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	testAssert := func(t testing.TB, got, want string) {
		t.Helper()
		if want != got {
			t.Errorf("Expected %q. Got %q", want, got)
		}
	}

	t.Run("Repeating single character 5 times", func(t *testing.T) {
		expected := "aaaaa"
		repeated := Repeat("a", 5)
		testAssert(t, repeated, expected)
	})

	t.Run("Repeating four character string 10 times", func(t *testing.T) {
		expected := "abcdabcdabcdabcdabcdabcdabcdabcdabcdabcd"
		repeated := Repeat("abcd", 10)
		testAssert(t, repeated, expected)
	})
}
func ExampleRepeat() {
	exampleOutput := Repeat("y", 5)
	fmt.Println(exampleOutput)
	// Output: yyyyy
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
