package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Fatal("add(1, 2) should be 3")
	}
}
