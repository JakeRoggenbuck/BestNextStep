package main

import (
	"testing"
)

func TestHash(t *testing.T) {
	input := "test_password"

	hash1, _ := HashPassword(input)
	first := CheckPasswordHash(input, hash1)

	hash2, _ := HashPassword(input)
	second := CheckPasswordHash(input, hash2)

	if !first {
		t.Errorf("hash %q was not valid", hash1)
	}

	if !second {
		t.Errorf("hash %q was not valid", hash2)
	}

	if hash1 == hash2 {
		t.Errorf("hashes where the same")
	}
}

