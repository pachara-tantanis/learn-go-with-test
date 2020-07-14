package main

import "testing"

func Test_helloWorld(t *testing.T) {
	got := Hello()
	want := "Hello World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}