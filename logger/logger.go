package logger

type Scope interface {
	Contain(Scope) bool
}

type Writer interface {
	Write([]byte) error
}

type Logger struct {
	scope  Scope
	writer Writer
}

// NewLogger returns a new logger.
func NewLogger(scope Scope) *Logger {
	return &Logger{
		scope:  scope,
		writer: nil,
	}
}

func (log *Logger) SetWriter(writer Writer) {
	log.writer = writer
}

func (log *Logger) Write(scope Scope, p []byte) error {
	// Ignore
	if !log.scope.Contain(scope) {
		return nil
	}
	return log.writer.Write(p)
}
