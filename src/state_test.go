package main

import (
	"testing"

	"github.com/satori/go.uuid"
)

func TestUpdateState(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := uuid.NewV4().String()
		if getState(id) != "unknown" {
			t.Error("invalid initial status")
		}
		updateState(id, "new")
		if getState(id) != "new" {
			t.Error("status failed to update")
		}
		updateStateWithDebug(id, "another update", "new2")
		if getState(id) != "new2" {
			t.Error("status failed to update")
		}
	}
}
