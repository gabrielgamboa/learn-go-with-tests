package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat a character 5 times when no repeatTimes is provided", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"

		assertIfRepeatedIsDifferentFromExpected(t, repeated, expected)
	})

	t.Run("repeat a character 'repeatedTimes' times when repeatTimes is provided", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"

		assertIfRepeatedIsDifferentFromExpected(t, repeated, expected)
	})
}

func assertIfRepeatedIsDifferentFromExpected(t testing.TB, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("expected %q but got %q", expected, got)
	}
}

func ExampleRepeat(t *testing.T) {
	response := Repeat("b", 5)
	fmt.Println(response)
	//Output: "bbbbb"
}

// go test -bench=. -benchmem
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 0)
	}
}
