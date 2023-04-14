package main

import "testing"

func TestCalculateRectangleArea(t *testing.T) {
	expected := 9
	actual := CalculateRectangleArea(4, 5)

	if expected != actual {
		t.Errorf("expected: %v, got: %v", expected, actual)
	}

}
