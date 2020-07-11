package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "string of text"}

	got := Search(dictionary, "test")
	want := "string of text"

	if got != want {
		t.Errorf("got %q instead of %q", got, want)
	}
}
