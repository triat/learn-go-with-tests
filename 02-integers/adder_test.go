package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(2, 4)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if expected != sum {
		t.Errorf("expected %d, got %d", expected, sum)
	}
}
