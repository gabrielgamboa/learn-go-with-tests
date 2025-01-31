package arraysandslices

import "testing"

func TestSum(t *testing.T) {
	numbers := [4]int{20, 10, 5, 5}

	response := Sum(numbers)
	expected := 40

	if response != expected {
		t.Errorf("expected %d but got %d. Numbers: %v", expected, response, numbers)
	}
}
