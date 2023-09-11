package logger

import (
	"os"

	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
}

type LoggerResult struct {
	Logger *zerolog.Logger
}

func NewLogger(p Params) *LoggerResult {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &LoggerResult{Logger: &logger}
}

func (lr *LoggerResult) Info(message string) {
	lr.Logger.Info().Msg(message)
}

func (lr *LoggerResult) Debug(message string) {
	lr.Logger.Debug().Msg(message)
}

func (lr *LoggerResult) Error(message string) {
	lr.Logger.Error().Msg(message)
}

func (lr *LoggerResult) Fatal(message string) {
	lr.Logger.Fatal().Msg(message)
}

func (lr *LoggerResult) LogRequest(message ...interface{}) {
	lr.Logger.Log().Fields(message).Msg("")
}
