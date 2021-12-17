package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	brewreadFile("test_files/Brewfile")
	output := Run("go", "sum")
	if output == "" {
		t.Errorf("Output was empty")
	}
	fmt.Println(t.TempDir())
}
