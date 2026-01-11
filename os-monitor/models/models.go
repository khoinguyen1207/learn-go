package models

import "context"

type Monitor interface {
	Name() string
	Check(ctx context.Context) string
}

type SystemStats struct {
	Name         string
	UsagePercent string
}
