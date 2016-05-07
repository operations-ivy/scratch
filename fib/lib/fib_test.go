package fib_test

import (
  "testing"

  "github.com/whitepatrick/scratch/fib/lib"
)

var fibTests = []struct {
  n        int // input
  expected int // exoected inout
}{
  {1,1},
  {2,1},
  {3,2},
  {4,3},
  {5,5},
  {6,8},
  {7,13},
}

func TestFib(t *testing.T) {
  for _, tt := range fibTests {
    actual := Fib(tt.n)
    if actual != tt.expected {
      t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
    }
  }
}
