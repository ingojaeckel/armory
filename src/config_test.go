package main

import "testing"

func TestUnitLoadConfiguration(t *testing.T) {
	if err := initConfiguration(); err != nil {
		t.Errorf("Failed to init configuration: %s", err.Error())
	}
	if conf.Base.Environment != "dev" {
		t.Error("Failed to initialize configuration correctly.")
	}
}
