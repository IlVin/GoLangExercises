package logex

import (
	"log"
	"os"
)

type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
)

type LogExtended struct {
	log.Logger
	lvl LogLevel
}

func NewLogExtended() *LogExtended {
	l := LogExtended{}
	l.SetOutput(os.Stderr)
	l.SetFlags(log.Ldate + log.Ltime)
	return &l
}

func (l *LogExtended) SetLogLevel(logLvl LogLevel) {
	l.lvl = logLvl
}

func (l *LogExtended) Infoln(msg string) {
	if l.lvl >= LogLevelInfo {
		l.Println("[INFO] ", msg)
	}
}

func (l *LogExtended) Warnln(msg string) {
	if l.lvl >= LogLevelWarning {
		l.Println("[WARN] ", msg)
	}
}

func (l *LogExtended) Errorln(msg string) {
	if l.lvl >= LogLevelError {
		l.Println("[ERROR] ", msg)
	}
}
