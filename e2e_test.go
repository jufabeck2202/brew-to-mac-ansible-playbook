package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	output := Run("go", "sum")
	if output == "" {
		t.Errorf("Output was empty")
	}
	fmt.Println(t.TempDir())
}
