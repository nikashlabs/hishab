package logger

import (
	"go.uber.org/zap"
)

type Zapper struct {
	log *zap.SugaredLogger
}

func NewZapLogger() (logger Logger, err error) {
	z, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	sugar := z.Sugar()
	return &Zapper{log: sugar}, nil
}

func (z *Zapper) Debug(msg string, keysAndValues ...any) {
	z.log.Debugw(msg, keysAndValues...)
}

func (z *Zapper) Info(msg string, keysAndValues ...any) {
	z.log.Infow(msg, keysAndValues...)
}

func (z *Zapper) Warn(msg string, keysAndValues ...any) {
	z.log.Warnw(msg, keysAndValues...)
}

func (z *Zapper) Error(msg string, keysAndValues ...any) {
	z.log.Errorw(msg, keysAndValues...)
}

func (z *Zapper) Fatal(msg string, keysAndValues ...any) {
	z.log.Fatalw(msg, keysAndValues...)
}

func (z *Zapper) Sync() error {
	return z.log.Sync()
}
