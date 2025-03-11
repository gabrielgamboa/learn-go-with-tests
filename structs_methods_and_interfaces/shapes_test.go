package structsmethodsandinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.10, 10.10}
	got := Perimeter(rectangle)
	want := 40.40
	assertMessage(t, got, want)
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %g but want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		want := 72.0

		checkArea(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793

		checkArea(t, circle, want)
	})
}

func assertMessage(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
