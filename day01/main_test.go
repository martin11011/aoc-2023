package main

import (
    "testing"
    _ "embed"
)

//go:embed test_input.txt
var testInput string

func Test_part1(t *testing.T) {
    result, err := part1(testInput)
    if err != nil {
        t.Errorf("Error: %v", err)
    }
    if result != 142 {
        t.Errorf("Sum was incorrect, got: %d, want: %d", result, 142)
    }
}

