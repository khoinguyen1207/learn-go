package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

const TRACE_ID_KEY = "trace_id"

var Log *zerolog.Logger

type LoggerConfig struct {
	Level      string // log level
	Filename   string // log file path
	MaxSize    int    // megabytes
	MaxBackups int    // number of backups
	MaxAge     int    //days
	Compress   bool   // disabled by default
	IsDev      string // development mode
}

func InitLogger(mode string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("❌ Error getting current working directory: %v", err)
	}
	filepath := filepath.Join(dir, "internal/logs/app.log")

	Log = NewLogger(LoggerConfig{
		Level:      "info",
		Filename:   filepath,
		MaxSize:    1, // megabytes
		MaxBackups: 5,
		MaxAge:     7, // days
		Compress:   true,
		IsDev:      mode,
	})
}

func NewLogger(cfg LoggerConfig) *zerolog.Logger {
	zerolog.TimeFieldFormat = time.RFC3339

	lvl, err := zerolog.ParseLevel(cfg.Level)
	if err != nil {
		lvl = zerolog.InfoLevel
	}

	fileWriter := &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	}
	var writer io.Writer = fileWriter
	if cfg.IsDev == "development" {
		// consoleWriter := PrettyJSONWriter{Writer: os.Stdout}
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		writer = io.MultiWriter(fileWriter, consoleWriter)
	}

	logger := zerolog.New(writer).Level(lvl).With().Timestamp().Logger()
	return &logger
}

func GetTraceId(ctx context.Context) string {
	if traceId, ok := ctx.Value(TRACE_ID_KEY).(string); ok {
		return traceId
	}
	return ""
}

// Custom writer to pretty-print JSON logs in the console
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
