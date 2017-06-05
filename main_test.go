package main

import (
	"os"
	"testing"
)

// Test: GetEnv
func TestGetEnv(t *testing.T) {
	os.Setenv("EXISTING_VAR", "1")

	v := GetEnv("EXISTING_VAR", "")

	if v != "1" {
		t.Error("Expected ok")
	}
}

// Test: GetEnv with fallback
func TestGetEnvFallback(t *testing.T) {
	v := GetEnv("NOT_EXISTING_VAR", "fallback")

	if v != "fallback" {
		t.Error("Expected \"fallback\" in fallback value")
	}
}
