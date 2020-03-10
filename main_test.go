package main

import (
	"os"
	"testing"
)

func TestConfigs_default(t *testing.T) {
	b := "https://www.example.com"
	os.Setenv("BACKEND", b)

	configure()
	if backend != b {
		t.Errorf("Expected %s, got %s", b, backend)
	}

	if port != "3000" {
		t.Errorf("Expected 3000, got %s", port)
	}

	if bind != "0.0.0.0" {
		t.Errorf("Expected 0.0.0.0, got %s", bind)
	}
}

func TestConfigs_configured(t *testing.T) {
	b := "https://www.example.com"
	os.Setenv("BACKEND", b)
	os.Setenv("PORT", "3333")
	os.Setenv("BIND", "localhost")
	configure()

	if port != "3333" {
		t.Errorf("Expected 3333, got %s", port)
	}

	if bind != "localhost" {
		t.Errorf("Expected localhost, got %s", bind)
	}
}
