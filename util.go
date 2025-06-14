package util

import (
	"os"
	"runtime"

	"golang.org/x/term"
)

func IsInteractive() bool {
	return term.IsTerminal(int(os.Stdin.Fd()))
}

func GetFuncName() string {
	name, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(name).Name()
}
