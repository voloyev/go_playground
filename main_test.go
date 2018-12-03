package main

import "testing"

var name = defFile{"bob", "body"}

func TestGetName(t *testing.T) {
	got := getName(&name)
	want := "body"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestChangeName(t *testing.T) {
	changed, _ := name.Write("new_body")
	got, _ := changed.Read()
	want := "new_body"

	if got != want {
		t.Errorf("Write() do not change name")
	}
}
