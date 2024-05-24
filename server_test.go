package main

import (
  "testing"
)

func TestTruthy(t *testing.T) {
  want := true
  result := Truthy()
  
  if want != result {
    t.Fatal("truthy should reture true, got false instead")
  }
} 
