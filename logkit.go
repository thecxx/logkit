package logkit

import (
	"io"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DefaultTimeLayout = "2006-01-02 15:04:05.999999999 -0700 MST"
)

const (
	LevelDebug = zapcore.DebugLevel
	LevelInfo  = zapcore.InfoLevel
	LevelWarn  = zapcore.WarnLevel
	LevelError = zapcore.ErrorLevel
	LevelPanic = zapcore.PanicLevel
)

type (
	Field = zapcore.Field
	Level = zapcore.Level
)

// K takes a key and a value, then return a field.
func K(key string, value interface{}) Field {
	return zap.Any(key, value)
}

type Option func(*LoggerConfig)

// WithLoggerEncoder sets a function to new an encoder.
func WithLoggerEncoder(fun EncoderNew) Option {
	return func(cfg *LoggerConfig) {
		cfg.encoderNew = fun
	}
}

// WithLoggerFieldKey sets a key name with specific key.
func WithLoggerFieldKey(key, name string) Option {
	return func(cfg *LoggerConfig) {
		cfg.encoderOpts = append(cfg.encoderOpts, WithEncoderFieldKey(key, name))
	}
}

// WithLoggerTimeEncoder sets a time encoder.
func WithLoggerTimeEncoder(loc *time.Location, layout string) Option {
	return func(cfg *LoggerConfig) {
		cfg.encoderOpts = append(cfg.encoderOpts, WithTimeEncoder(loc, layout))
	}
}

// WithLoggerCaller sets caller in log.
func WithLoggerCaller(skip int) Option {
	return func(cfg *LoggerConfig) {
		cfg.zloggerOpts = append(cfg.zloggerOpts, zap.AddCaller(), zap.AddCallerSkip(skip))
	}
}

var (
	defaultConfig = LoggerConfig{
		level:       LevelDebug,
		writer:      NewConsoleWriter(),
		encoderNew:  NewJSONEncoder,
		encoderOpts: []EncoderOption{WithTimeEncoder(time.Local, DefaultTimeLayout)},
		zloggerOpts: make([]zap.Option, 0),
	}
	log = (*Logger)(nil)
)

type LoggerConfig struct {
	level       Level
	writer      io.Writer
	encoderNew  EncoderNew
	encoderOpts []EncoderOption
	zloggerOpts []zap.Option
}

type Logger struct {
	zlog *zap.Logger
}

// NewLoggerWithOptions returns a logger.
func NewLogger(opts ...Option) *Logger {
	// Copy default logger config
	cfg := defaultConfig
	// Apply options
	for _, fun := range opts {
		fun(&cfg)
	}
	newEnc := cfg.encoderNew
	// New zap logger
	zlog := zap.New(
		zapcore.NewCore(
			newEnc(cfg.encoderOpts...),
			zapcore.AddSync(cfg.writer),
			cfg.level,
		),
	)
	if len(cfg.zloggerOpts) > 0 {
		zlog = zlog.WithOptions(cfg.zloggerOpts...)
	}
	return &Logger{zlog}
}

func (l *Logger) Debug(msg string, fields ...Field) {
	l.zlog.Debug(msg, fields...)
}

func Debug(msg string, fields ...Field) {
	log.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.zlog.Info(msg, fields...)
}

func Info(msg string, fields ...Field) {
	log.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.zlog.Warn(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	log.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.zlog.Error(msg, fields...)
}

func Error(msg string, fields ...Field) {
	log.Error(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...Field) {
	l.zlog.Panic(msg, fields...)
}

func Panic(msg string, fields ...Field) {
	log.Panic(msg, fields...)
}

// Init initializes default logger.
func Init(opts ...Option) {
	log = NewLogger(opts...)
}
