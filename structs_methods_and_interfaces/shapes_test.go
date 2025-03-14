package structsmethodsandinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.10, 10.10}
	got := Perimeter(rectangle)
	want := 40.40
	assertMessage(t, got, want)
}

func TestArea(t *testing.T) {
	areaTests := []struct { //table driven tests => multiple cases to test one interface
		name      string
		shape     Shape
		shapeArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 16}, shapeArea: 192.0},
		{name: "Circle", shape: Circle{10}, shapeArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, shapeArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			want := tt.shapeArea

			if got != want {
				t.Errorf("%v: got %g want %g", tt.shape, got, want)
			}
		})
	}
}

func assertMessage(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
