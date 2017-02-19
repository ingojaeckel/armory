package main

import "testing"

func TestUnitCreateNew(t *testing.T) {
	if NewTest().ID == NewTest().ID {
		t.Error("test ids must not match")
	}
}
