package main

import "testing"

func TestExecuteGatling(t *testing.T) {
	err := executeGatling("com.example.Loadtest")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestUnitExecute(t *testing.T) {
	err := execute([]string{}, "/bin/ls", "-la")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestUnitExecute2(t *testing.T) {
	err := execute([]string{"KEY=value"}, "/bin/ls", "-la", "/tmp/")
	if err != nil {
		t.Fatal(err.Error())
	}
}
