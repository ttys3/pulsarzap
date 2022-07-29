package pulsarzap

import (
	"github.com/apache/pulsar-client-go/pulsar/log"
	"go.uber.org/zap"
)

// zapWrapper implements Logger interface based on zap SugaredLogger
type zapWrapper struct {
	l *zap.SugaredLogger
}

func NewDefault() log.Logger {
	logger, _ := zap.NewProduction(zap.WithCaller(true))
	return New(logger.Sugar())
}

func New(lgr *zap.SugaredLogger) log.Logger {
	lgr = lgr.Desugar().WithOptions(zap.AddCallerSkip(1)).Sugar()
	return &zapWrapper{
		l: lgr,
	}
}

func (l *zapWrapper) SubLogger(fs log.Fields) log.Logger {
	// pulsar use this like: `logger.SubLogger(log.Fields{"serviceURL": serviceURL})`
	// ref https://github.com/apache/pulsar-client-go/blob/bd19458b32ff89206c135cc647336e690e99c32f/pulsar/internal/rpc_client.go#L83
	return &zapWrapper{
		l: l.l.With(pulsarFieldsToKVSlice(fs)...),
	}
}

func (l *zapWrapper) WithFields(fs log.Fields) log.Entry {
	return &zapWrapper{
		l: l.l.With(pulsarFieldsToKVSlice(fs)...),
	}
}

func (l *zapWrapper) WithField(name string, value interface{}) log.Entry {
	return &zapWrapper{
		l: l.l.With(name, value),
	}
}

func (l *zapWrapper) WithError(err error) log.Entry {
	return &zapWrapper{
		l: l.l.With("err", err),
	}
}

func (l *zapWrapper) Debug(args ...interface{}) {
	l.l.Debug(args...)
}

func (l *zapWrapper) Info(args ...interface{}) {
	l.l.Info(args...)
}

func (l *zapWrapper) Warn(args ...interface{}) {
	l.l.Warn(args...)
}

func (l *zapWrapper) Error(args ...interface{}) {
	l.l.Error(args...)
}

func (l *zapWrapper) Debugf(format string, args ...interface{}) {
	l.l.Debugf(format, args...)
}

func (l *zapWrapper) Infof(format string, args ...interface{}) {
	l.l.Infof(format, args...)
}

func (l *zapWrapper) Warnf(format string, args ...interface{}) {
	l.l.Warnf(format, args...)
}

func (l *zapWrapper) Errorf(format string, args ...interface{}) {
	l.l.Errorf(format, args...)
}

func (l *zapWrapper) With(fs log.Fields) log.Entry {
	return &zapWrapper{
		l: l.l.With(pulsarFieldsToKVSlice(fs)...),
	}
}
