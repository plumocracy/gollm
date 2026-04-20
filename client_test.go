package gollm

import "testing"

// This is just a base test to ensure that things are running properly in github actions.
func TestOneNeqZero(t *testing.T) {
	value := 1
	if value != value-1 {
		t.Errorf("Value %d == %d???", value, value-1)
	}
}
