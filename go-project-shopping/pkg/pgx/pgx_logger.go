package pgx

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/tracelog"
	"github.com/rs/zerolog"
)

type PgxZeroLogTracer struct {
	Logger         zerolog.Logger
	SlowQueryLimit time.Duration
}

func (t *PgxZeroLogTracer) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]any) {
	sql, _ := data["sql"].(string)
	// args, _ := data["args"].([]any)
	// duration, _ := data["time"].(time.Duration)

	// log.Println("⚠️ Data:", data)
	// log.Println("⚠️ SQL:", sql)
	// log.Println("⚠️ Args:", args)
	// log.Println("⚠️ Duration:", duration)

	baseLogger := t.Logger.With().
		Str("sql", sql)
	// 	Interface("args", args).
	// 	Dur("duration", duration)

	logger := baseLogger.Logger()
	logger.Info().Str("event", "Query").Msg("Executed SQL")

	// if msg == "Query" && duration > t.SlowQueryLimit {
	// 	logger.Warn().Str("event", "Slow Query").Msg("Slow SQL Query")
	// 	return
	// }

	// if msg == "Query" {
	// 	logger.Info().Str("event", "Query").Msg("Executed SQL")
	// 	return
	// }

}
