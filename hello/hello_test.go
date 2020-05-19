package hello

import (
  "testing"
)

func TestHello(t *testing.T) {
  result := GetHello("HI! Guest!")
  expext := "HI! Guest! Hello, World!"
  if result != expext {
    t.Error("\nactual： ", result, "\nexpected： ", expext)
  }

  t.Log("TestHello End.")
}
