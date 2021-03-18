package logkit

import (
	"time"

	"go.uber.org/zap/zapcore"
)

const (
	EncoderFieldKeyMessage    = "message"
	EncoderFieldKeyLevel      = "level"
	EncoderFieldKeyTime       = "time"
	EncoderFieldKeyCaller     = "caller"
	EncoderFieldKeyFunction   = "function"
	EncoderFieldKeyStacktrace = "stacktrace"
)

var (
	defaultEncoderConfig = EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    "function",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
)

type (
	Encoder       = zapcore.Encoder
	EncoderConfig = zapcore.EncoderConfig
)

type EncoderNew func(...EncoderOption) Encoder
type EncoderOption func(*EncoderConfig)

// WithEncoderFieldKey sets a new name for key.
func WithEncoderFieldKey(key, name string) EncoderOption {
	return func(cfg *EncoderConfig) {
		switch key {
		case EncoderFieldKeyTime:
			cfg.TimeKey = name
		case EncoderFieldKeyLevel:
			cfg.LevelKey = name
		case EncoderFieldKeyMessage:
			cfg.MessageKey = name
		case EncoderFieldKeyCaller:
			cfg.CallerKey = name
		case EncoderFieldKeyFunction:
			cfg.FunctionKey = name
		case EncoderFieldKeyStacktrace:
			cfg.StacktraceKey = name
		}
	}
}

// WithTimeEncoder sets a new time encoder.
func WithTimeEncoder(loc *time.Location, layout string) EncoderOption {
	return func(cfg *EncoderConfig) {
		cfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.In(loc).Format(layout))
		}
	}
}

// NewJSONEncoder takes some options and then return an json encoder.
func NewJSONEncoder(opts ...EncoderOption) Encoder {
	// Copy default encoder config
	cfg := defaultEncoderConfig
	// Apply options
	for _, fun := range opts {
		fun(&cfg)
	}
	return zapcore.NewJSONEncoder(cfg)
}

// NewConsoleEncoder takes some options and then return an console encoder.
func NewConsoleEncoder(opts ...EncoderOption) Encoder {
	// Copy default encoder config
	cfg := defaultEncoderConfig
	// Apply options
	for _, fun := range opts {
		fun(&cfg)
	}
	return zapcore.NewConsoleEncoder(cfg)
}
