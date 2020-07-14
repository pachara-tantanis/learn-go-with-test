package main

import "testing"

func Test_helloWorld(t *testing.T) {
	got := Hello("Yok")
	want := "Hello, Yok"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}