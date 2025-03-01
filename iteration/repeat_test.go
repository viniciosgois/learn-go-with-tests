package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	output := Repeat("a", 5)
	fmt.Println(output)
	// output = "aaaaa"
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestRepeat(t *testing.T) {
	t.Run("should repeat the given times", func(t *testing.T) {
		repeated := Repeat("a", 10)
		expected := "aaaaaaaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

}
