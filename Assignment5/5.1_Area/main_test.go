// Testing
package main

import (
	"math"
	"testing"
)

const tolerance = 0.00001

func TestCost_Square(t *testing.T) {
	// Testcase1
	s1 := Square{10.0}
	expected1 := 10.0 * 10.0 * 0.02
	output1 := Cost(s1)
	if math.Abs(output1-expected1) > tolerance {
		t.Errorf("Expected: %f, got: %f for testcase: %v", expected1, output1, s1)
	}

	// Testcase2
	s2 := Square{5.0}
	expected2 := 5.0 * 5.0 * 0.02
	output2 := Cost(s2)
	if math.Abs(output2-expected2) > tolerance {
		t.Errorf("Expected: %f, got: %f for testcase: %v", expected2, output2, s2)
	}
}

func TestCost_Rectangle(t *testing.T) {
	// Testcase1
	r1 := Rectangle{4.0, 5.0}
	expected1 := 4.0 * 5.0 * 0.42
	output1 := Cost(r1)
	if math.Abs(output1-expected1) > tolerance {
		t.Errorf("Expected: %f, got: %f for testcase: %v", expected1, output1, r1)
	}

	// Testcase2
	r2 := Rectangle{10.0, 2.0}
	expected2 := 10.0 * 2.0 * 0.42
	output2 := Cost(r2)
	if math.Abs(output2-expected2) > tolerance {
		t.Errorf("Expected: %f, got: %f for testcase: %v", expected2, output2, r2)
	}
}

func TestCost_Circle(t *testing.T) {
	// Testcase1
	c1 := Circle{3.0}
	expected1 := 3.0 * 3.0 * math.Pi * 1.20
	output1 := Cost(c1)
	if math.Abs(output1-expected1) > tolerance {
		t.Errorf("Expected: %f, got: %f for testcase: %v", expected1, output1, c1)
	}

	// Testcase2
	c2 := Circle{7.0}
	expected2 := 7.0 * 7.0 * math.Pi * 1.20
	output2 := Cost(c2)
	if math.Abs(output2-expected2) > tolerance {
		t.Errorf("Expected: %.20f, got: %.20f for testcase: %v", expected2, output2, c2)
	}
}
