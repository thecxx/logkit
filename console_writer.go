package logkit

import (
	"io"
	"os"
)

type ConsoleWriter struct {
	out io.Writer
}

// NewConsoleWriter returns a console writer.
func NewConsoleWriter() io.Writer {
	return &ConsoleWriter{os.Stdout}
}

func (c *ConsoleWriter) Write(p []byte) (n int, err error) {
	return c.out.Write(p)
}
