package util

import (
	"fmt"
	"io"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	LogPanic = "[P]"
	LogFatal = "[F]"
	LogErr   = "[E]"
	LogWarn  = "[W]"
	LogInfo  = "[I]"
	LogDebug = "[D]"
	LogTrace = "[T]"
)

var (
	Log              *log.Logger
	logTag           string
	logCallback      UiLogCallback    = nil
	progressCallback ProgressCallback = nil
)

type UiLogCallback func(string, string)
type ProgressCallback func(bool)

func SetUiLogCallback(callback UiLogCallback) {
	logCallback = callback
}

func SetProgressCallback(callback ProgressCallback) {
	progressCallback = callback
}

func SetProgress(en bool) {
	if progressCallback != nil {
		progressCallback(en)
	}
}

type customFormatter struct {
	log.TextFormatter
}

func (f *customFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := time.Now().Format("15:04:05")

	switch entry.Level {
	case log.TraceLevel:
		logTag = LogTrace
		break
	case log.DebugLevel:
		logTag = LogDebug
		break
	case log.InfoLevel:
		logTag = LogInfo
		break
	case log.WarnLevel:
		logTag = LogWarn
		break
	case log.ErrorLevel:
		logTag = LogErr
		break
	case log.FatalLevel:
		logTag = LogFatal
		break
	default:
		logTag = LogPanic
		break
	}

	logMessage := []byte(timestamp + " " + logTag + " " + entry.Message + "\n")
	return logMessage, nil
}

type uiWriter struct{}

func (uw *uiWriter) Write(p []byte) (n int, err error) {

	if logCallback != nil {
		logCallback(logTag, string(p))
	}
	return len(p), nil
}

func LogInit(logFile *os.File) {
	Log = log.New()
	Log.SetFormatter(&customFormatter{})
	Log.SetLevel(log.DebugLevel)

	var mw io.Writer
	writers := []io.Writer{&uiWriter{}}

	if logFile != nil {
		writers = append(writers, logFile)
	}
	if IsInteractive() {
		writers = append(writers, os.Stdout)
	}
	mw = io.MultiWriter(writers...)
	Log.SetOutput(mw)
}

func DumpBuffer(data []byte, colum int, hex bool) {
	out := ""
	for i, d := range data {
		if hex {
			out = out + fmt.Sprintf("0x%02x,", d)
		} else {
			out = out + fmt.Sprintf("%d,", d)
		}
		if ((i + 1) % colum) == 0 {
			Log.Println(out)
			out = ""
		}
	}
	if out != "" {
		Log.Println(out)
	}
}
