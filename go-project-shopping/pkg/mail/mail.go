package mail

import (
	"context"
	"project-shopping/internal/config"
	"project-shopping/pkg/logger"
	"time"

	"github.com/rs/zerolog"
)

type MailMessage struct {
	From     Address
	To       []Address
	Cc       []Address
	Subject  string
	BodyText string
	BodyHTML string
}

type Address struct {
	Name  string
	Email string
}

type SendResult struct {
	MessageID string
	Provider  string
}

type MailConfig struct {
	FromAddress string
	MaxRetries  int
}

// ======================== MailService ========================
type mailService struct {
	provider MailProvider
	config   MailConfig
	log      *zerolog.Logger
}

func NewMailService(config *config.Config, provider MailProvider, logger *zerolog.Logger) MailService {
	mailConfig := MailConfig{
		FromAddress: config.Mail.FromAddress,
		MaxRetries:  3,
	}

	return &mailService{
		provider: provider,
		config:   mailConfig,
		log:      logger,
	}
}

func (ms *mailService) SendMail(ctx context.Context, msg *MailMessage) error {
	traceID := logger.GetTraceId(ctx)
	start := time.Now()

	var lastErr error
	for attempt := 1; attempt <= ms.config.MaxRetries; attempt++ {
		startAttempt := time.Now()
		_, err := ms.provider.SendMail(ctx, msg)
		if err == nil {
			ms.log.Info().
				Str("trace_id", traceID).
				Dur("duration", time.Since(startAttempt)).
				Str("operation", "send_mail").
				Interface("to", msg.To).
				Interface("subject", msg.Subject).
				Msg("Email sent successfully")
			return nil
		}

		lastErr = err
		ms.log.Warn().
			Err(err).
			Str("trace_id", traceID).
			Dur("duration", time.Since(startAttempt)).
			Str("operation", "send_mail").
			Int("attempt", attempt).
			Msg("Failed to send email")

		time.Sleep(time.Duration(attempt) * time.Second) // Exponential backoff
	}

	ms.log.Error().
		Err(lastErr).
		Str("trace_id", traceID).
		Dur("duration", time.Since(start)).
		Str("operation", "send_mail").
		Msg("Failed to send email after max retries")

	return lastErr
}
