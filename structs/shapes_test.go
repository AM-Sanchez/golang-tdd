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

	// Use an anonymous struct to perform table-driven tests.
	// This allows a developer to add a new type with an Area()
	// function and be able to easily add the type to the test.
	// This strategy also permits the easy addition of extra Area() test
	// cases as needed.
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 6, Height: 8}, want: 48.0},
		{shape: Circle{Radius: 8}, want: 201.06192982974676},
		{shape: Triangle{Base: 9, Height: 3}, want: 13.5},
	}

	for _, testItem := range areaTests {
		got := testItem.shape.Area()
		if testItem.want != got {
			t.Errorf("got: %g | want %g", got, testItem.want)
		}
	}
}
