package main

import "testing"

func TestUnitGatlingConfig(t *testing.T) {
	initConfiguration()
	if conf.Base.GatlingRoot == "" {
		t.Error("Gatling Root is not set")
	}
	if conf.GatlingBinary == "/bin/gatling.sh" {
		t.Error("Gatling binary directly points at /. Is your GatlingRoot set?")
	}
	if conf.ResultsFolder == "/results/testname" {
		t.Error("Gatling results folder points at /. Is your GatlingRoot set?")
	}
}

func TestUnitLoadConfiguration(t *testing.T) {
	initConfiguration()
	if conf.Base.GatlingRoot == "" {
		t.Errorf("Failed to init configuration")
	}
	if conf.Base.Environment != "dev" {
		t.Error("Failed to initialize configuration correctly.")
	}
}
