package iface

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{11.0, 20.0}
	got := Perimeter(rectangle)
	want := 62.0
	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 12.0}

		got := rectangle.Area()
		want := 120.0

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}

		got := circle.Area()
		want := math.Pi * math.Pow(10.0, 2)

		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	})
}

func TestShape(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
		name  string
	}{
		{Rectangle{10.0, 12.0}, 120.0, "rectangles"},
		{Circle{10.0}, math.Pi * 10.0 * 10.0, "circles"},
		{Triangle{10.0, 10.0}, 50.0, "triangles"},
	}
	for _, areaTest := range areaTests {
		t.Run(areaTest.name, func(t *testing.T) {
			got := areaTest.shape.Area()
			want := areaTest.want

			if got != want {
				t.Errorf("got %g, want %g", got, want)
			}
		})
	}
}
