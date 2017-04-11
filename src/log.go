package main

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/sebest/logrusly"
)

var loggingInitialized = false
var lastFlush = time.Now()
var logger *logrus.Logger
var hook *logrusly.LogglyHook

func Log(msg string, args ...interface{}) {
	LogWithFlush(fmt.Sprintf(msg, args...), shouldFlush())
}

func shouldFlush() bool {
	if !conf.Base.LogglyEnabled {
		return false
	}

	flushExpiryTime := time.Now().Add(-conf.FlushDuration)
	if !lastFlush.Before(flushExpiryTime) {
		return false
	}

	lastFlush = time.Now()
	return true
}

func LogWithFlush(msg string, flush bool) {
	if !conf.Base.LogglyEnabled {
		fmt.Println(msg)
		return
	}

	initializeLogging()
	logger.WithFields(logrus.Fields{"env": conf.Base.Environment}).Info(msg)
	if flush {
		hook.Flush()
	}
}

func initializeLogging() {
	if !conf.Base.LogglyEnabled {
		return
	}
	if !loggingInitialized {
		logger = logrus.New()
		hook = logrusly.NewLogglyHook(conf.Base.LogglyToken, "https://logs-01.loggly.com/bulk/", logrus.InfoLevel, "go", "logrus")
		logger.Hooks.Add(hook)

		loggingInitialized = true
	}
}
