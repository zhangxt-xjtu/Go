package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	epxected := 4

	if sum != epxected {
		t.Errorf("expected '%d' but got '%d'", epxected, sum)
	}
}
