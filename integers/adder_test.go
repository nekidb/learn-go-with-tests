package integers

import "testing"
import "fmt"

func TestAdder(t *testing.T) {

	t.Run("testing adder", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected: %d, got: %d", expected, sum)
		}
	})
}

func ExampleAdd() {
	sum := Add(-1, 1001)
	fmt.Println(sum)
	// Output: 1000
}
