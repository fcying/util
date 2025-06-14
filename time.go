package util

import (
	"sync"
	"time"
)

var timeStart sync.Map

func PreciseDelayMs(dly uint32) {
	start := time.Now()
	for {
		if (time.Since(start) / time.Millisecond) >= time.Duration(dly) {
			break
		}
	}
}

func PreciseDelayUs(dly uint32) {
	start := time.Now()
	for {
		if (time.Since(start) / time.Microsecond) >= time.Duration(dly) {
			break
		}
	}
}

func TimeStart(key string) {
	timeStart.Store(key, time.Now())
}

func TimeDelay(key string) uint32 {
	value, ok := timeStart.Load(key)
	if !ok {
		return 0
	}
	startTime := value.(time.Time)
	delay := uint32(time.Since(startTime) / time.Millisecond)
	return delay
}

func TimeEnd(key string, text string) uint32 {
	value, ok := timeStart.Load(key)
	if !ok {
		return 0
	}
	startTime := value.(time.Time)
	delay := uint32(time.Since(startTime) / time.Millisecond)
	if text != "NULL" {
		Log.Infof("%s %d ms%s", key, delay, text)
	}
	timeStart.Delete(key)
	return delay
}
