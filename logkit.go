package logkit

import (
	"github.com/thecxx/logkit/logger"
)

type Scope struct {
	level uint32
}

// NewScope returns a scope.
func NewScope(level uint32) logger.Scope {
	return &Scope{level}
}

func (s *Scope) Contain(scope logger.Scope) bool {
	v, ok := scope.(*Scope)
	if !ok {
		return false
	}
	return s.level <= v.level
}

type Field struct {
	key   string
	value interface{}
}

func Debug(message string) {

}

func Info(message string) {

}

func Warn(message string) {

}

func Error(message string) {

}
