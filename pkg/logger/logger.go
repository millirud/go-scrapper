package logger

import (
	"io"

	"github.com/rs/zerolog"
)

func New(level string, w io.Writer) zerolog.Logger {
	lvl, err := zerolog.ParseLevel(level)

	if err != nil {
		panic(err)
	}

	return zerolog.New(w).With().
		Timestamp().
		Logger().
		Level(lvl)
}