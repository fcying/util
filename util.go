package util

import (
	"os"
	"runtime"
	"time"

	"golang.org/x/term"
)

var timeStart map[string]time.Time

func init() {
	timeStart = make(map[string]time.Time)
}

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
	timeStart[key] = time.Now()
}

func TimeDelay(key string) uint32 {
	delay := uint32(time.Since(timeStart[key]) / 1000000)
	return delay
}

func TimeEnd(key string, text string) uint32 {
	delay := uint32(time.Since(timeStart[key]) / 1000000)
	if text != "NULL" {
		Log.Infof("%s %d ms%s", key, delay, text)
	}
	delete(timeStart, key)

	return delay
}

func IsInteractive() bool {
	return term.IsTerminal(int(os.Stdin.Fd()))
}

func GetFuncName() string {
	name, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(name).Name()
}
