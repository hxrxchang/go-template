package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	result := square(2)
	if result != 4 {
		t.Fatal("test failed")
	}
}
