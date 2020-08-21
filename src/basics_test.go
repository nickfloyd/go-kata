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
	{1, "The value is SMALL"},
	{7, "The value is MEDIUM"},
	{100, "The value is LARGE"},
}

func TestIfStatements(t *testing.T) {
	for _, tt := range conditionalTests {
		actual := ifStatements(tt.n)
		if actual != tt.expected {
			t.Errorf("If Statements(%d): expected: %s, actual: %s", tt.n, tt.expected, actual)
		}
	}
}

func TestSwitchStatements(t *testing.T) {
	for _, tt := range conditionalTests {
		actual := switchStatements(tt.n)
		if actual != tt.expected {
			t.Errorf("Switch Statement(%d): expected: %s, actual: %s", tt.n, tt.expected, actual)
		}
	}
}
