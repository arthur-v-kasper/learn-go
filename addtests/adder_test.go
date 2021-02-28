package interation

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	sumAdd := ExampleAdd(2, 2, 1)

	if sum != expected {
		t.Errorf("expected (sum) '%d' but got '%d'", expected, sum)
	}

	if sumAdd != expected {
		t.Errorf("expected (sumAdd) '%d' but got '%d'", expected, sumAdd)
	}
}

func Add(x, y int) int {
	return x + y
}

func ExampleAdd(x, y, z int) int {
	return x + y + z
}
