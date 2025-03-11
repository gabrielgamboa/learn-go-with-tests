package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 4 values", func(t *testing.T) {
		numbers := []int{20, 10, 5, 5}

		response := Sum(numbers)
		expected := 40

		if response != expected {
			t.Errorf("expected %d but got %d. Numbers: %v", expected, response, numbers)
		}
	})

	t.Run("collection of any values", func(t *testing.T) {
		numbers := []int{20, 10, 5, 5, 10, 50}

		response := Sum(numbers)
		expected := 100

		if response != expected {
			t.Errorf("expected %d but got %d. Numbers: %v", expected, response, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) { //DeepEqual is not type safe
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9})
		want := []int{5, 9}

		if !reflect.DeepEqual(got, want) { //DeepEqual is not type safe
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
