package mail

import (
	"context"
	"fmt"
	"project-shopping/internal/config"
	"project-shopping/pkg/logger"
	"project-shopping/pkg/template"
	"time"

	"github.com/rs/zerolog"
)

type MailMessage struct {
	To       []Address
	Cc       []Address
	Subject  string
	BodyText string
	BodyHTML string
}

type MailMessageTemplate struct {
	To           []Address
	Cc           []Address
	Subject      string
	TemplateName string
	Data         any
}

type Address struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email"`
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
	template template.TemplateService
}

func NewMailService(config *config.Config, provider MailProvider, logger *zerolog.Logger, tmpl template.TemplateService) MailService {
	mailConfig := MailConfig{
		FromAddress: config.MailFromAddress,
		MaxRetries:  3,
	}

	return &mailService{
		provider: provider,
		config:   mailConfig,
		log:      logger,
		template: tmpl,
	}
}

func (ms *mailService) SendWithTemplate(
	ctx context.Context,
	msg *MailMessageTemplate,
) error {
	html, err := ms.template.Render(msg.TemplateName, msg.Data)
	if err != nil {
		return fmt.Errorf("mail: render template failed: %w", err)
	}

	return ms.SendMail(ctx, &MailMessage{
		To:       msg.To,
		Cc:       msg.Cc,
		Subject:  msg.Subject,
		BodyHTML: html,
	})
}
func (ms *mailService) SendMail(ctx context.Context, msg *MailMessage) error {
	traceID := logger.GetTraceId(ctx)
	start := time.Now()

	var lastErr error
	for attempt := 1; attempt <= ms.config.MaxRetries; attempt++ {
		startAttempt := time.Now()
		result, err := ms.provider.SendMail(ctx, msg)
		if err == nil {
			ms.log.Info().
				Str("trace_id", traceID).
				Dur("duration", time.Since(startAttempt)).
				Str("operation", "send_mail").
				Str("provider", ms.provider.Name()).
				Interface("to", msg.To).
				Interface("subject", msg.Subject).
				Interface("messageId", result.MessageID).
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
			Str("provider", ms.provider.Name()).
			Msg("Failed to send email")

		time.Sleep(time.Duration(attempt) * time.Second) // Exponential backoff
	}

	ms.log.Error().
		Err(lastErr).
		Str("trace_id", traceID).
		Dur("duration", time.Since(start)).
		Str("operation", "send_mail").
		Str("provider", ms.provider.Name()).
		Msg("Failed to send email after max retries")

	return lastErr
}
