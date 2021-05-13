package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("mozzie")
	want := "Hello, mozzie"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
