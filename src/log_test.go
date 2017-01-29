package main

import "testing"

func TestUnitLogging(t *testing.T) {
	Log("automated %s %d", "1", 2)
	LogWithFlush("automated tests.", true)
}
