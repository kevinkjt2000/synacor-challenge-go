package main

import "testing"

func TestStuff(t *testing.T) {
	got := stuff()
	if got != 0 {
		t.Errorf("stuff() did not return 0, instead got: %d", got)
	}
}
