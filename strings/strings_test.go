package strings

import (
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	count := strings.Count("strawberry", "r")
	expected := 3

	if count != expected {
		t.Errorf("expected %d but got %d", expected, count)
	}
}

func TestCut(t *testing.T) {
	before, after, _ := strings.Cut("strawberry", "r")

	result := before + after
	expected := "stawberry"

	if result != expected {
		t.Errorf("expected %q but got %q", expected, result)
	}
}
