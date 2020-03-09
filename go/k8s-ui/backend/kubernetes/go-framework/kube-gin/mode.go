package kube_gin

import (
	"os"
	"io"
)

const EnvGinMode = "GIN_MODE"

const (
	debugCode = iota
	releaseCode
	testCode
)

const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)

var ginMode = debugCode
var modeName = DebugMode

var DefaultWriter io.Writer = os.Stdout

func init() {
	mode := os.Getenv(EnvGinMode)
	SetMode(mode)
}

// SetMode sets gin mode according to input string.
func SetMode(value string) {
	switch value {
	case DebugMode, "":
		ginMode = debugCode
	case ReleaseMode:
		ginMode = releaseCode
	case TestMode:
		ginMode = testCode
	default:
		panic("gin mode unknown: " + value)
	}
	if value == "" {
		value = DebugMode
	}
	modeName = value
}
// Mode returns currently gin mode.
func Mode() string {
	return modeName
}
