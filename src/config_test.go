package main

import "testing"

func TestUnitLoadConfiguration(t *testing.T) {
	initConfiguration()
	if conf.Base.GatlingRoot == "" {
		t.Errorf("Failed to init configuration")
	}
	if conf.Base.Environment != "dev" {
		t.Error("Failed to initialize configuration correctly.")
	}
}
