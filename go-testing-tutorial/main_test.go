package main

import "testing"

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 { // Checks if Calculate == 4 if it doesnt it executes the code
		t.Error("Expected 2 + 2 to equal 4") // Throws an error
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{99999, 100001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}

}
