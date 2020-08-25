package main

import (
	"testing"
)

func TestNewRandomXingamento(t *testing.T) {
	result := NewRandomXingamento()

	if result == "" {
		t.Errorf("Offense is null, can't generate phrases")
	}
}