package logger

import "go.uber.org/zap"

type Logger interface {
	With(args ...any) Logger
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

type logger struct {
	base *zap.SugaredLogger
}

func New() (Logger, error) {
	l, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	l = l.WithOptions(zap.AddCallerSkip(1))

	return &logger{base: l.Sugar()}, nil
}

func (l *logger) With(args ...any) Logger {
	return &logger{base: l.base.With(args...)}
}

func (l *logger) Debug(msg string, args ...any) {
	l.base.Debugw(msg, args)
}

func (l *logger) Info(msg string, args ...any) {
	l.base.Infow(msg, args)
}

func (l *logger) Error(msg string, args ...any) {
	l.base.Errorw(msg, args...)
}
