package logger

import (
	"sync"

	"go.uber.org/zap"
)

type Zapper struct {
	log *zap.SugaredLogger
}

var _ Logger = (*Zapper)(nil)

var (
	instance Logger
	once     sync.Once
)

func GetInstance() Logger {
	once.Do(func() {
		z, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}
		instance = &Zapper{log: z.Sugar()}
	})
	return instance
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

var Log = GetInstance()
