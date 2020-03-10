package main

import (
	"testing"
	"os"
)

func TestConfigs_default(t *testing.T) {
	b := "https://www.example.com"
	os.Setenv("BACKEND", b)

	configure()
	if backend != b {
		t.Errorf("Expected %s, got %s", b, backend)
	}
}