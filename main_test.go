package main

import "testing"

func TestAdder(t *testing.T) {
  res1 := Adder(90, 34.56)
  if res1 != 124.56 {
    t.Errorf("Expected 124.56, got %v\n", res1)
  }
}
