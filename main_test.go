package main

import "testing"

func TestHello(t *testing.T) {
	res := hello()
	expected := "Hello Jenkins"
    if res != expected {
       t.Errorf("Hello was incorrect, got: %s, want: %s.", res, expected)
    }
}