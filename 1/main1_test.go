package main

import (
	"crypto/sha256"
	"testing"
)

func TestHashFunction(t *testing.T) {
	input := "42 0o52 2A 3.14 Golang true (1+2i)"
	expectedHash := sha256.Sum256([]byte("42 4go-202424 42 3.14 Golang true (1+2i)"))

	runeSlice := []rune(input)
	saltedRunes := append(runeSlice[:len(runeSlice)/2], append([]rune("go-2024"), runeSlice[len(runeSlice)/2:]...)...)
	hash := sha256.Sum256([]byte(string(saltedRunes)))

	if hash != expectedHash {
		t.Errorf("Ожидался hash %x, получен %x", expectedHash, hash)
	}
}
