package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Havy")
	got := buffer.String()
	want := "Hello, Havy"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
