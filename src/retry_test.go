package main

import (
	"errors"
	"testing"
	"time"
)

var globalCounter = 0

func TestUnitRetry(t *testing.T) {
	beforeMs := time.Now().UnixNano() / 1000 / 1000

	// attempt1, sleep 100ms, attempt2, sleep 200ms, attempt 3, sleep 250, attempt 4 -> end
	// total sleep time: 550ms
	if err := retry(true, 100, 250, 10, foo); err != nil {
		t.Errorf("Should have passed with retries")
	}
	afterMs := time.Now().UnixNano() / 1000 / 1000
	bufferMs := 50
	if afterMs-beforeMs > int64(100+200+250+bufferMs) {
		t.Errorf("Sleep time too long: %d", (afterMs - beforeMs))
	}
}

func foo() (bool, error) {
	globalCounter++

	if globalCounter == 4 {
		return true, nil
	}
	return false, errors.New("fake error")
}
