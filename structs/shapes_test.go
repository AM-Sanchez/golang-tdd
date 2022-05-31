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
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle Area Test", shape: Rectangle{Width: 6, Height: 8}, hasArea: 48.0},
		{name: "Circle Area Test", shape: Circle{Radius: 8}, hasArea: 201.06192982974676},
		{name: "Right-Triangle Area Test", shape: Triangle{Base: 9, Height: 3}, hasArea: 13.5},
	}

	for _, testItem := range areaTests {
		// Use the name field of the anonymous struct as the name of the test.
		t.Run(testItem.name, func(t *testing.T) {
			got := testItem.shape.Area()
			if testItem.hasArea != got {
				t.Errorf("%#v got: %g | has area: %g", testItem.shape, got, testItem.hasArea)
			}
		})
	}
}
