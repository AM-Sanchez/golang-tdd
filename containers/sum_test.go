package containers

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 2},
		[]int{1, 2, 1, 1, 1},
		[]int{3, 3, 3, 3, 3, 3, 3, 3, 6},
		[]int{20, 10, 30, 5, 5, 8, 4, 2, 6, 1, 2, 3, 4})
	want := []int{5, 6, 30, 100}

	// Use reflect.DeepEqual() instead of using for loops.
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTail(t *testing.T) {

	checkSums := func(t testing.TB, want, got []int) {
		t.Helper()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("Wanted %d, got %d", want, got)
		}
	}

	t.Run("All slices passed in are size > 1", func(t *testing.T) {
		want := []int{5, 10, 15}
		got := SumAllTails([]int{5, 2, 1, 2},
			[]int{20, 2, 2, 2, 2, 2},
			[]int{8, 3, 4, 1, 2, 5})
		checkSums(t, want, got)
	})

	t.Run("One slice passed is size of 1 and another is empty", func(t *testing.T) {
		want := []int{0, 0, 15}
		got := SumAllTails([]int{}, []int{20}, []int{8, 3, 4, 1, 2, 5})
		checkSums(t, want, got)
	})

	t.Run("Pass in a single empty slice", func(t *testing.T) {
		want := []int{0}
		got := SumAllTails([]int{})
		checkSums(t, want, got)
	})
}
