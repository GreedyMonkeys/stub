package model

import "testing"

func TestAdd(t *testing.T) {
	result := add(1, 2)
	want := 3
	if result != want {
		t.Errorf("got %d want %d", result, want)
	}
}
