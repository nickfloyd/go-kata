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
