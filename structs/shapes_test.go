package structs

import (
	"testing"
)

func TestPerimieter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f | want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(5.0, 10.0)
	want := 50.0

	if got != want {
		t.Errorf("got: %.2f | want %.2f", got, want)
	}
}
