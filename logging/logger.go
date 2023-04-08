package logging

import "fmt"

type Level = int8

type Logger struct {
	level Level
}

const (
	LevelDebug Level = iota
	LevelInfo
)

func NewLogger(level Level) Logger {
	return Logger{level: level}
}

func (l Logger) log(level Level, msg string, params ...any) {
	if level >= l.level {
		fmt.Printf(msg+"\n", params...)
	}
}

func (l Logger) LogDebug(msg string, params ...any) {
	l.log(LevelDebug, msg, params...)
}

func (l Logger) LogInfo(msg string, params ...any) {
	l.log(LevelInfo, msg, params...)
}
