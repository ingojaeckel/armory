package main

import "time"

func retry(debugOutput bool, initialSleepTimeMillis int, maxSleepTimeMillis int, maxAttempts int, fn func() (bool, error)) error {
	sleepTimeMilliseconds := initialSleepTimeMillis
	var lastErr error

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		isDone, err := fn()

		if debugOutput {
			Log("Attempt #%d", attempt)
		}
		if err != nil {
			Log("Caught error: %s", err.Error())
		}
		if isDone {
			return nil
		}
		if debugOutput {
			Log("Sleeping for %d ms", sleepTimeMilliseconds)
		}
		time.Sleep(time.Duration(sleepTimeMilliseconds) * time.Millisecond)

		// Increase sleep time for next attempt
		if sleepTimeMilliseconds*2 > maxSleepTimeMillis {
			sleepTimeMilliseconds = maxSleepTimeMillis
		} else {
			sleepTimeMilliseconds *= 2
		}
		lastErr = err
	}

	return lastErr
}
