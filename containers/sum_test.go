package containers

import "testing"

func TestSum(t *testing.T) {
	assertForTestSlice := func(t testing.TB, want, got int, given []int) {
		t.Helper()
		if want != got {
			t.Errorf("Wanted %d, got %d, given %v", want, got, given)
		}
	}

	t.Run("Slice of four numbers", func(t *testing.T) {
		want := 10
		vals := []int{1, 2, 3, 4}
		got := Sum(vals)
		assertForTestSlice(t, want, got, vals)
	})
}
