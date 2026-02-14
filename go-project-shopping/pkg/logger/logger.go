package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

const TraceIdKey = "trace_id"

type LoggerConfig struct {
	Level      string // log level
	Filename   string // log file path
	MaxSize    int    // megabytes
	MaxBackups int    // number of backups
	MaxAge     int    //days
	Compress   bool   // disabled by default
	IsDev      string // development mode
}

func NewLoggerConfig(cfg LoggerConfig) *zerolog.Logger {
	zerolog.TimeFieldFormat = time.RFC3339

	lvl, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		lvl = zerolog.InfoLevel
	}
	// zerolog.SetGlobalLevel(lvl)

	var writer io.Writer
	if cfg.IsDev == "development" {
		writer = PrettyJSONWriter{Writer: os.Stdout}
	} else {
		writer = &lumberjack.Logger{
			Filename:   cfg.Filename,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
		}
	}

	logger := zerolog.New(writer).Level(lvl).With().Timestamp().Logger()

	return &logger
}

type PrettyJSONWriter struct {
	Writer io.Writer
}

func (w PrettyJSONWriter) Write(p []byte) (n int, err error) {
	var prettyJSON bytes.Buffer

	err = json.Indent(&prettyJSON, p, "", "  ")
	if err != nil {
		return w.Writer.Write(p)
	}

	return w.Writer.Write(prettyJSON.Bytes())
}

func GetTraceId(ctx context.Context) string {
	if traceId, ok := ctx.Value(TraceIdKey).(string); ok {
		return traceId
	}
	return ""
}
