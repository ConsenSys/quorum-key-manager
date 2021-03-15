package log

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	entry     *logrus.Entry
	component string
	ctx       context.Context
	cfg       *Config
}

func NewDefaultLogger() *Logger {
	logger := logrus.New()
	return &Logger{
		entry: logrus.NewEntry(logger),
	}
}

func NewLogger(cfg *Config) *Logger {
	if cfg == nil {
		return NewDefaultLogger()
	}

	logger := logrus.New()
	logger.Level = cfg.Level.logrusLvl()
	if cfg.Timestamp {
		logger.SetFormatter(&logrus.TextFormatter{
			PadLevelText:     true,
			FullTimestamp:    true,
			DisableTimestamp: false,
		})
	}

	return &Logger{
		entry: logrus.NewEntry(logger),
		cfg:   cfg,
	}
}

func (l Logger) WithContext(ctx context.Context) *Logger {
	l.ctx = ctx
	l.entry = l.entry.WithContext(ctx)
	return &l
}

func (l Logger) WithError(err error) *Logger {
	l.entry = l.entry.WithError(err)
	return &l
}

func (l Logger) SetComponent(c string) *Logger {
	l.component = c
	return &l
}

func (l *Logger) Component() string {
	return l.component
}

func (l Logger) WithFields(fields logrus.Fields) *Logger {
	l.entry = l.entry.WithFields(fields)
	return &l
}

func (l Logger) WithField(key string, value interface{}) *Logger {
	return l.WithFields(logrus.Fields{key: value})
}

func (l *Logger) Debug(args ...interface{}) {
	l.entry.WithFields(contextLogFields(l.ctx)).Debug(l.args(args)...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.entry.WithFields(contextLogFields(l.ctx)).Debug(l.format(format), args)
}

func (l *Logger) Info(args ...interface{}) {
	l.entry.WithFields(contextLogFields(l.ctx)).Info(l.args(args)...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.entry.WithFields(contextLogFields(l.ctx)).Infof(l.format(format), args)
}

func (l *Logger) Error(args ...interface{}) {
	l.entry.WithFields(contextLogFields(l.ctx)).Error(l.args(args)...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.entry.WithFields(contextLogFields(l.ctx)).Errorf(l.format(format), args)
}

func (l *Logger) Warn(args ...interface{}) {
	l.entry.WithFields(contextLogFields(l.ctx)).Warn(l.args(args)...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.entry.WithFields(contextLogFields(l.ctx)).Warnf(l.format(format), args)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.entry.WithFields(contextLogFields(l.ctx)).Fatal(l.args(args)...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.entry.WithFields(contextLogFields(l.ctx)).Fatalf(l.format(format), args)
}

func (l *Logger) Trace(args ...interface{}) {
	l.entry.WithFields(contextLogFields(l.ctx)).Trace(l.args(args)...)
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	l.entry.WithFields(contextLogFields(l.ctx)).Trace(l.format(format), args)
}

func (l *Logger) Write(p []byte) (n int, err error) {
	switch l.entry.Level {
	case logrus.ErrorLevel:
		l.Error(string(p))
	case logrus.WarnLevel:
		l.Warn(string(p))
	case logrus.InfoLevel:
		l.Info(string(p))
	case logrus.DebugLevel:
		l.Debug(string(p))
	case logrus.TraceLevel:
		l.Trace(string(p))
	default:
		l.Info(string(p))
	}

	return 0, nil
}

func (l *Logger) args(args []interface{}) []interface{} {
	return append([]interface{}{l.format("")}, args...)
}

func (l *Logger) format(msg string) string {
	if l.component == "" {
		return msg
	}

	return fmt.Sprintf("(%s): %s", l.component, msg)
}
