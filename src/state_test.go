package main

import (
	"fmt"
	"testing"

	"github.com/satori/go.uuid"
)

func BenchmarkUpdateState(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		id := uuid.NewV4().String()
		newState := fmt.Sprintf("new state for id=%s", id)
		b.StartTimer()
		updateState(id, newState)
	}
}
