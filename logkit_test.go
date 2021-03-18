package logkit

import (
	"testing"
)

func TestInfo(t *testing.T) {
	Init(WithLoggerCaller(2))
	Info("test info")
}
