package logging

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Log Levels
const (
	LevelDebug = 1 << iota
	LevelInfo
	LevelWarning
	LevelError
)

const (
	debugColor   = "\u001b[34m"
	infoColor    = "\u001b[36m"
	warningColor = "\u001b[33m"
	errorColor   = "\u001b[31m"
	clearColor   = "\u001b[0m"

	timeFormat = "2006-01-02 15:04:05"
)

func init() {
	defaultLogger = New()
}

var defaultLogger Logger

// Logger is a custom interface used to write to an output
// in a standard format.
type Logger interface {
	// Debug writes to the logger's output using fmt.Sprintf,
	// with the debug color scheme. The logger must have the
	// LevelDebug log level set.
	Debug(format string, args ...interface{})

	// Info writes to the logger's output using fmt.Sprintf,
	// with the info color scheme. The logger must have either
	// LevelDebug or LevelInfo set.
	Info(format string, args ...interface{})

	// Warning writes to the logger's output using fmt.Sprintf,
	// with the warning color scheme. The logger must have either
	// LevelDebug, LevelInfo or LevelWarning set.
	Warning(format string, args ...interface{})

	// Error writes to the logger's output using fmt.Sprintf,
	// with the warning color scheme. The logger must have either
	// LevelDebug, LevelInfo, LevelWarning or LevelError set.
	Error(format string, args ...interface{})

	// Sets the logger's output log level.
	SetLevel(lvl int) Logger

	// SetName sets the logger's output name.
	//
	// An example output is: [2006-01-02 15:04:05][{name}]: Hello World
	SetName(name string) Logger

	// SetOutput sets the output writer of the Logger.
	SetOutput(output io.Writer) Logger
}

type logger struct {
	output io.Writer
	lvl    int
	name   string
}

// New returns a new instance of the Logger interface with the given output.
func New() Logger {
	return &logger{
		output: os.Stderr,
		lvl:    LevelInfo,
	}
}

func (l *logger) Debug(format string, args ...interface{}) {
	l.print(format, args, debugColor, LevelDebug)
}

func (l *logger) Info(format string, args ...interface{}) {
	l.print(format, args, infoColor, LevelInfo)
}

func (l *logger) Warning(format string, args ...interface{}) {
	l.print(format, args, warningColor, LevelWarning)
}

func (l *logger) Error(format string, args ...interface{}) {
	l.print(format, args, errorColor, LevelError)
}

func (l *logger) print(format string, args []interface{}, color string, lvl int) {
	if l.lvl > lvl {
		return
	}

	var name string
	switch lvl {
	case LevelDebug:
		name = "DEBUG"
		break
	case LevelInfo:
		name = "INFO"
		break
	case LevelWarning:
		name = "WARNING"
		break
	case LevelError:
		name = "ERROR"
		break
	}

	if l.name != "" {
		name = l.name
	}

	prefix := fmt.Sprintf("[%s][%s%s%s]: ", time.Now().Format(timeFormat), color, name, clearColor)

	txt := prefix + fmt.Sprintf(format, args...) + "\n"
	l.output.Write([]byte(txt))
}

func (l *logger) SetLevel(lvl int) Logger {
	l.lvl = lvl
	return l
}

func (l *logger) SetName(name string) Logger {
	l.name = name
	return l
}

func (l *logger) SetOutput(output io.Writer) Logger {
	l.output = output
	return l
}

// Debug writes to os.Stderr using fmt.Sprintf, with the debug color
// scheme. The default logger must have either the LevelDebug log level set.
func Debug(format string, args ...interface{}) {
	defaultLogger.Debug(format, args...)
}

// Info writes to os.Stderr using fmt.Sprintf, with the info color
// scheme. The default logger must have either the LevelDebug or LevelInfo log level set.
func Info(format string, args ...interface{}) {
	defaultLogger.Info(format, args...)
}

// Warning writes to os.Stderr using fmt.Sprintf, with the warning color
// scheme. The default logger must have either the LevelDebug, LevelInfo or
// LevelWarning log level set.
func Warning(format string, args ...interface{}) {
	defaultLogger.Warning(format, args...)
}

// Error writes to os.Stderr using fmt.Sprintf, with the error color
// scheme. The default logger must have either the LevelDebug, LevelInfo,
// LevelWarning or LevelError log level set.
func Error(format string, args ...interface{}) {
	defaultLogger.Error(format, args...)
}

// SetLevel sets the logging level of the default logger.
func SetLevel(lvl int) Logger {
	return defaultLogger.SetLevel(lvl)
}

// SetName sets the log name of the default logger.
func SetName(name string) Logger {
	return defaultLogger.SetName(name)
}

// SetOutput sets the output write of the default logger.
func SetOutput(output io.Writer) Logger {
	return defaultLogger.SetOutput(output)
}
