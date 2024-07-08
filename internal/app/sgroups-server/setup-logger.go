package server

import (
	"github.com/wildberries-tech/sgroups/v2/internal/app"

	"github.com/H-BF/corlib/logger"
	"github.com/pkg/errors"
)

func SetupLogger() error {
	ctx := app.Context()
	_, err := LoggerLevel.Value(ctx, LoggerLevel.OptSink(func(v string) error {
		var l logger.LogLevel
		if e := l.UnmarshalText([]byte(v)); e != nil {
			return errors.Wrapf(e, "recognize '%s' logger level from config", v)
		}
		logger.SetLevel(l)
		return nil
	}))
	return err
}
