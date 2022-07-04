package main

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestGlob(t *testing.T) {
	matches, err := filepath.Glob("*.go")
	if err != nil {
		panic(err)
	}
	for _, match := range matches {
		fmt.Printf("match: %v\n", match)
	}
}
