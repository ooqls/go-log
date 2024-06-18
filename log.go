package log

import (
	"go.uber.org/zap"
)

var level zap.AtomicLevel = zap.NewAtomicLevelAt(zap.DebugLevel)
var l *zap.Logger

func init() {

	cfg := zap.Config{
		Level:            level,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		Encoding:         "json",
	}

	var err error
	l, err = cfg.Build()
	if err != nil {
		panic(err)
	}

}

func NewLogger(name string) *zap.Logger {
	newL := l.With(zap.String("source", name))
	return newL
}
