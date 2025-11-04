package marketdata

import (
	"log"
	"os"
)

// Logger is a reusable logging utility
type Logger struct {
	file   *os.File
	logger *log.Logger
}

// NewLogger creates a new logger instance that logs to both console and file
func NewLogger(filename string) (*Logger, error) {
	// Open or create log file
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	// Create MultiWriter for file + console output
	l := log.New(file, "", log.LstdFlags|log.Lmicroseconds)

	return &Logger{
		file:   file,
		logger: l,
	}, nil
}

// Info logs an informational message
func (l *Logger) Info(format string, v ...interface{}) {
	l.logger.Printf("[INFO] "+format, v...)
}

// Error logs an error message
func (l *Logger) Error(format string, v ...interface{}) {
	l.logger.Printf("[ERROR] "+format, v...)
}

// Close closes the log file
func (l *Logger) Close() {
	if l.file != nil {
		l.file.Close()
	}
}
