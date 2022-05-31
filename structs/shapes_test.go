package structs

import (
	"testing"
)

type Shape interface {
	Area() float64
}

func TestPerimieter(t *testing.T) {
	rect := Rectangle{10, 10}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f | want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	// Create helper interface to allow for a Shape to be passed in
	// This should either be a Rectangle or Circle type and error on any other.
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got: %g | want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{6.0, 8.0}
		want := 48.0
		checkArea(t, rect, want)
	})

	t.Run("circles", func(t *testing.T) {
		circ := Circle{8.0}
		want := 201.06192982974676
		checkArea(t, circ, want)
	})

}
