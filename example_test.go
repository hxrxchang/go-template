package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	result := example(2)
	if result != 4 {
		t.Fatal("test failed")
	}
}
