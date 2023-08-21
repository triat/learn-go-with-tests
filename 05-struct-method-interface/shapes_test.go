package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if want != got {
		t.Errorf("expected %.2f got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12.0, Height: 6.0}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()

		if tt.want != got {
			t.Errorf("%#v expected %g got %g", tt.shape, tt.want, got)
		}
	}
}
