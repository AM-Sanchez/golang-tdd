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
	// want := []int{5, 6, 30, 100}
	// given := [][]int{{1, 2, 3},
	// 	{1, 2, 1, 1, 1},
	// 	{3, 3, 3, 3, 3, 3, 3, 3, 6},
	// 	{20, 10, 30, 5, 5, 8, 4, 2, 6, 1, 2, 3, 4}}
	// got := SumAll(given)

	// success := false
	// for index, val := range want {
	// 	gotSum := 0
	// 	for _, givenVal := range given[index] {
	// 		gotSum += givenVal
	// 	}

	// 	if gotSum != val {
	// 		success = false
	// 	}
	// }
	// if !success {
	// 	t.Errorf("Wanted %d, got %d, given %v", want, got, given)
	// }
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
