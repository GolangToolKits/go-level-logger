package levellogger

import "log"

const (
	//OffLevel OffLevel
	OffLevel = 0
	//InfoLevel InfoLevel
	InfoLevel = 1
	//DebugLevel DebugLevel
	DebugLevel = 2
	//AllLevel AllLevel
	AllLevel = 3
)

// Log Log
type Log interface {
	Info(val ...any)
	Debug(val ...any)
	Error(val ...any)
	SetLogLevel(level int)
}

// Logger Logger
type Logger struct {
	LogLevel int
}

// New New
func (l *Logger) New() Log {
	return l
}

// Info Info
func (l *Logger) Info(val ...any) {
	if l.LogLevel == InfoLevel || l.LogLevel == AllLevel {
		log.Println("INFO: ", val)
	}
}

// Debug Debug
func (l *Logger) Debug(val ...any) {
	if l.LogLevel == DebugLevel || l.LogLevel == AllLevel {
		log.Println("DEBUG: ", val)
	}
}

// Error Error
func (l *Logger) Error(val ...any) {
	log.Println("ERROR: ", val)
}

// SetLogLevel SetLogLevel
func (l *Logger) SetLogLevel(level int) {
	l.LogLevel = level
}

//go mod init github.com/GolangToolKits/go-level-logger
