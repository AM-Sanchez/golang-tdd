package containers

import "testing"

func TestSum(t *testing.T) {
	want := 10
	vals := [4]int{1, 2, 3, 4}
	got := Sum(vals)
	if want != got {
		t.Errorf("Wanted %d, got %d given %v", want, got, vals)
	}
}
