package dora

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/NoneBorder/dora/zlogwriter"

	"github.com/rs/zerolog"
)

func init() {
	zerolog.TimeFieldFormat = time.RFC3339
}

// Logger is the global logger.
var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

// AutoWriter return logger which writer decider by adapterName
// @logLevel should be string as debug,info,warn,error,fatal,panic
func NewLogWithWriter(adapterName, writerConfig string, logLevel string, bufferSize ...int) zerolog.Logger {
	bufferSize = append(bufferSize, 0)

	logger := Logger
	switch adapterName {
	case "console":
		logger = TextWriter(nil)

	case "file":
		logger = FileWriter(writerConfig, bufferSize[0])
	}

	l, _ := zerolog.ParseLevel(logLevel)
	if l == zerolog.NoLevel {
		l = zerolog.DebugLevel
	}
	return logger.Level(l)
}

// TextWriter set return logger to zerolog.ConsoleWriter
func TextWriter(w io.Writer, noColor ...bool) zerolog.Logger {
	if w == nil {
		w = os.Stderr
	}
	noColor = append(noColor, false)
	return zerolog.New(zerolog.ConsoleWriter{Out: w, NoColor: noColor[0]}).With().Timestamp().Logger()
}

// FileWriter return logger to zlog/writer.FileWriter
func FileWriter(jsonConfig string, bufferSize ...int) zerolog.Logger {
	bufferSize = append(bufferSize, 0)
	w := zlogwriter.NewFileWriter()
	if err := w.Init(jsonConfig); err != nil {
		Logger.Fatal().Err(err).Msg("init zlogwriter.FileWriter failed")
		return Logger
	}
	return zerolog.New(w).With().Timestamp().Logger()
}

// Output duplicates the global logger and sets w as its output.
func Output(w io.Writer) zerolog.Logger {
	return Logger.Output(w)
}

// With creates a child logger with the field added to its context.
func With() zerolog.Context {
	return Logger.With()
}

// Level creates a child logger with the minimum accepted level set to level.
func Level(level zerolog.Level) zerolog.Logger {
	return Logger.Level(level)
}

// Sample returns a logger with the s sampler.
func Sample(s zerolog.Sampler) zerolog.Logger {
	return Logger.Sample(s)
}

// Hook returns a logger with the h Hook.
func Hook(h zerolog.Hook) zerolog.Logger {
	return Logger.Hook(h)
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func Debug() *zerolog.Event {
	return Logger.Debug()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func Info() *zerolog.Event {
	return Logger.Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func Warn() *zerolog.Event {
	return Logger.Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func Error() *zerolog.Event {
	return Logger.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func Fatal() *zerolog.Event {
	return Logger.Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func Panic() *zerolog.Event {
	return Logger.Panic()
}

// WithLevel starts a new message with level.
//
// You must call Msg on the returned event in order to send the event.
func WithLevel(level zerolog.Level) *zerolog.Event {
	return Logger.WithLevel(level)
}

// Log starts a new message with no level. Setting zerolog.GlobalLevel to
// zerolog.Disabled will still disable events produced by this method.
//
// You must call Msg on the returned event in order to send the event.
func Log() *zerolog.Event {
	return Logger.Log()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	Logger.Print(v...)
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	Logger.Printf(format, v...)
}

// Ctx returns the Logger associated with the ctx. If no logger
// is associated, a disabled logger is returned.
func Ctx(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}
