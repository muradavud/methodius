package logger

import (
	"github.com/rs/zerolog"
	"os"
	"strings"
)

type ILogger interface {
	Pretty(message string, args ...interface{})
	Debug(message string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
	Fatal(message string, args ...interface{})
}

type Logger struct {
	logger       zerolog.Logger
	prettyLogger zerolog.Logger
}

var _ ILogger = (*Logger)(nil)
var Log ILogger

func New(level string) *Logger {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)

	skipFrameCount := 3
	logger := zerolog.
		New(os.Stdout).
		With().
		Timestamp().
		CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).
		Logger()

	return &Logger{
		logger:       logger,
		prettyLogger: logger.Output(zerolog.ConsoleWriter{Out: os.Stderr}),
	}
}

func (l *Logger) Pretty(message string, args ...interface{}) {
	if len(args) == 0 {
		l.prettyLogger.Info().Msg(message)
	} else {
		l.prettyLogger.Info().Msgf(message, args...)
	}
}

func (l *Logger) Debug(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Debug().Msg(message)
	} else {
		l.logger.Debug().Msgf(message, args...)
	}
}

func (l *Logger) Info(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Info().Msg(message)
	} else {
		l.logger.Info().Msgf(message, args...)
	}
}

func (l *Logger) Warn(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Warn().Msg(message)
	} else {
		l.logger.Warn().Msgf(message, args...)
	}
}

func (l *Logger) Error(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Error().Msg(message)
	} else {
		l.logger.Error().Msgf(message, args...)
	}
}

func (l *Logger) Fatal(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Fatal().Msg(message)
	} else {
		l.logger.Fatal().Msgf(message, args...)
	}
}
