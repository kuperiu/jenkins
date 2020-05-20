package main

import "testing"

func TestHello(t *testing.T) {
    res := hello()
    if res != "hello world" {
       t.Errorf("Hello was incorrect, got: %s, want: %s.", res, "hello world")
    }
}