package main

import "testing"

func TestUnitCp(t *testing.T) {
	if err := cp([]pair{pair{"cp.go", "cp.txt"}}); err != nil {
		t.Error(err.Error())
	}
	if err := execute([]string{}, "rm", "cp.txt"); err != nil {
		t.Error(err.Error())
	}
}
