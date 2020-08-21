package basics

import (
	"testing"
)

func TestAssignments(t *testing.T) {

	result := assignments(10, 5)
	if result != 7 {
		t.Errorf("Mean was incorrect, got: %d, want: %d.", result, 7)
	}
}

type conditionalTestData struct {
	n        int
	expected string
}

var conditionalTests = []conditionalTestData {
	{1, "The value is SMALL"}, {7, "The value is MEDIUM"}, {100, "The value is LARGE"},
}

func TestConditionals(t *testing.T) {
	for _, tt := range conditionalTests {
			actual := conditionals(tt.n)
			if actual != tt.expected {
					t.Errorf("conditionals(%d): expected: %s, actual: %s", tt.n, tt.expected, actual)
			}
	}
}