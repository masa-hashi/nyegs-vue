package hello

import (
  "testing"
)

func TestHelloJapanease(t *testing.T) {
  result := GetHello("こんにちわ")
  expext := "こんにちわ Hello, World!"
  if result != expext {
    t.Error("\nactual： ", result, "\nexpected： ", expext)
  }

  t.Log("TestHello End.")
}

func TestHelloEnglish(t *testing.T) {
  result := GetHello("Hello!")
  expext := "Hello! Hello, World!"
  if result != expext {
    t.Error("\nactual： ", result, "\nexpected： ", expext)
  }

  t.Log("TestHello End.")
}

func TestHello(t *testing.T) {
  result := GetHello("")
  expext := "Hello! Hello, World!"
  if result != expext {
    t.Error("\nactual： ", result, "\nexpected： ", expext)
  }

  t.Log("TestHello End.")
}