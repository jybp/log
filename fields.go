// Package log provides a structured logs abstraction with context-aware
// logging. It output logs in the logfmt format.
// Storing log fields in the context allows to avoid passing loggers around for that
// sole purpose.
//
//	ctx := log.CtxWithField(ctx, "key", "value")
//	log.InfoC(ctx, "message")
//
// It is currently a simple wrapper around logrus, and can be used as a drop-in replacement.
package log

import (
	"context"

	"github.com/sirupsen/logrus"
)

// Logger interface to avoid leadking entryCtx.
type Logger interface {
	WithField(k string, v interface{}) Logger
	WithFields(fields Fields) Logger
	WithError(err error) Logger

	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})

	DebugC(ctx context.Context, args ...interface{})
	InfoC(ctx context.Context, args ...interface{})
	WarnC(ctx context.Context, args ...interface{})
	ErrorC(ctx context.Context, args ...interface{})
	FatalC(ctx context.Context, args ...interface{})

	DebugfC(ctx context.Context, format string, args ...interface{})
	InfofC(ctx context.Context, format string, args ...interface{})
	WarnfC(ctx context.Context, format string, args ...interface{})
	ErrorfC(ctx context.Context, format string, args ...interface{})
	FatalfC(ctx context.Context, format string, args ...interface{})
}

// Context-aware functions.

func DebugC(ctx context.Context, args ...interface{}) {
	fields, _ := ctx.Value(fieldsCtxKey).(Fields)
	logrus.WithFields(logrus.Fields(fields)).Debug(args...)
}

func InfoC(ctx context.Context, args ...interface{}) {
	fields, _ := ctx.Value(fieldsCtxKey).(Fields)
	logrus.WithFields(logrus.Fields(fields)).Info(args...)
}

func WarnC(ctx context.Context, args ...interface{}) {
	fields, _ := ctx.Value(fieldsCtxKey).(Fields)
	logrus.WithFields(logrus.Fields(fields)).Warn(args...)
}

func ErrorC(ctx context.Context, args ...interface{}) {
	fields, _ := ctx.Value(fieldsCtxKey).(Fields)
	logrus.WithFields(logrus.Fields(fields)).Error(args...)
}

func FatalC(ctx context.Context, args ...interface{}) {
	fields, _ := ctx.Value(fieldsCtxKey).(Fields)
	logrus.WithFields(logrus.Fields(fields)).Fatal(args...)
}

func DebugfC(ctx context.Context, format string, args ...interface{}) {
	fields, _ := ctx.Value(fieldsCtxKey).(Fields)
	logrus.WithFields(logrus.Fields(fields)).Debugf(format, args...)
}

func InfofC(ctx context.Context, format string, args ...interface{}) {
	fields, _ := ctx.Value(fieldsCtxKey).(Fields)
	logrus.WithFields(logrus.Fields(fields)).Infof(format, args...)
}

func WarnfC(ctx context.Context, format string, args ...interface{}) {
	fields, _ := ctx.Value(fieldsCtxKey).(Fields)
	logrus.WithFields(logrus.Fields(fields)).Warnf(format, args...)
}

func ErrorfC(ctx context.Context, format string, args ...interface{}) {
	fields, _ := ctx.Value(fieldsCtxKey).(Fields)
	logrus.WithFields(logrus.Fields(fields)).Errorf(format, args...)
}

func FatalfC(ctx context.Context, format string, args ...interface{}) {
	fields, _ := ctx.Value(fieldsCtxKey).(Fields)
	logrus.WithFields(logrus.Fields(fields)).Fatalf(format, args...)
}

type entryCtx struct {
	*logrus.Entry
}

func (e entryCtx) WithField(k string, v interface{}) Logger {
	return entryCtx{e.Entry.WithField(k, v)}
}

func (e entryCtx) WithFields(fields Fields) Logger {
	return entryCtx{e.Entry.WithFields(logrus.Fields(fields))}
}

func (e entryCtx) WithError(err error) Logger {
	return e.WithField(logrus.ErrorKey, err)
}

func (e entryCtx) DebugC(ctx context.Context, args ...interface{}) {
	fields, _ := ctx.Value(fieldsCtxKey).(Fields)
	e.Entry.WithFields(logrus.Fields(fields)).Debug(args...)
}

func (e entryCtx) InfoC(ctx context.Context, args ...interface{}) {
	fields, _ := ctx.Value(fieldsCtxKey).(Fields)
	e.Entry.WithFields(logrus.Fields(fields)).Info(args...)
}

func (e entryCtx) WarnC(ctx context.Context, args ...interface{}) {
	fields, _ := ctx.Value(fieldsCtxKey).(Fields)
	e.Entry.WithFields(logrus.Fields(fields)).Warn(args...)
}

func (e entryCtx) ErrorC(ctx context.Context, args ...interface{}) {
	fields, _ := ctx.Value(fieldsCtxKey).(Fields)
	e.Entry.WithFields(logrus.Fields(fields)).Error(args...)
}

func (e entryCtx) FatalC(ctx context.Context, args ...interface{}) {
	fields, _ := ctx.Value(fieldsCtxKey).(Fields)
	e.Entry.WithFields(logrus.Fields(fields)).Fatal(args...)
}

func (e entryCtx) DebugfC(ctx context.Context, format string, args ...interface{}) {
	fields, _ := ctx.Value(fieldsCtxKey).(Fields)
	e.Entry.WithFields(logrus.Fields(fields)).Debugf(format, args...)
}

func (e entryCtx) InfofC(ctx context.Context, format string, args ...interface{}) {
	fields, _ := ctx.Value(fieldsCtxKey).(Fields)
	e.Entry.WithFields(logrus.Fields(fields)).Infof(format, args...)
}

func (e entryCtx) WarnfC(ctx context.Context, format string, args ...interface{}) {
	fields, _ := ctx.Value(fieldsCtxKey).(Fields)
	e.Entry.WithFields(logrus.Fields(fields)).Warnf(format, args...)
}

func (e entryCtx) ErrorfC(ctx context.Context, format string, args ...interface{}) {
	fields, _ := ctx.Value(fieldsCtxKey).(Fields)
	e.Entry.WithFields(logrus.Fields(fields)).Errorf(format, args...)
}

func (e entryCtx) FatalfC(ctx context.Context, format string, args ...interface{}) {
	fields, _ := ctx.Value(fieldsCtxKey).(Fields)
	e.Entry.WithFields(logrus.Fields(fields)).Fatalf(format, args...)
}

// Directly uses the logrus default logger (logfmt).
// https://github.com/sirupsen/logrus/blob/master/logger.go#L87-L96

// To allow drop-in replacement.
type Fields map[string]interface{}

func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func WithField(k string, v interface{}) Logger {
	return entryCtx{logrus.WithField(k, v)}
}

func WithFields(fields Fields) Logger {
	return entryCtx{logrus.WithFields(logrus.Fields(fields))}
}

func WithError(err error) Logger {
	return WithField(logrus.ErrorKey, err)
}

// context-aware log fields
type ctxKey struct{}

var fieldsCtxKey ctxKey = struct{}{}

func CtxWithField(ctx context.Context, k string, v interface{}) context.Context {
	fields, ok := ctx.Value(fieldsCtxKey).(Fields)
	if !ok {
		fields = make(Fields)
	}
	fields[k] = v
	return context.WithValue(ctx, fieldsCtxKey, fields)
}

func CtxWithFields(ctx context.Context, fields Fields) context.Context {
	existing, ok := ctx.Value(fieldsCtxKey).(Fields)
	if !ok {
		existing = make(Fields)
	}
	for k, v := range fields {
		existing[k] = v
	}
	return context.WithValue(ctx, fieldsCtxKey, existing)
}

func FromCtx(ctx context.Context) Logger {
	fields, _ := ctx.Value(fieldsCtxKey).(logrus.Fields)
	return entryCtx{logrus.WithFields(fields)}
}
