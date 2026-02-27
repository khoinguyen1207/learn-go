package mail

import (
	"context"
	"net/http"
	"time"
)

type MailtrapProviderConfig struct {
	MailSender string
	NameSender string
	APIKey     string
	BaseURL    string
	Timeout    time.Duration
}

type MailtrapProvider struct {
	client *http.Client
	cfg    *MailtrapProviderConfig
}

func NewMailtrapProvider(cfg *MailtrapProviderConfig) (*MailtrapProvider, error) {
	return &MailtrapProvider{
		client: &http.Client{Timeout: cfg.Timeout},
		cfg:    cfg,
	}, nil
}

func (p *MailtrapProvider) Name() string {
	return "mailtrap"
}

func (p *MailtrapProvider) SendMail(ctx context.Context, msg *MailMessage) (*SendResult, error) {
	return &SendResult{
		MessageID: "mailtrap-message-id-placeholder",
		Provider:  p.Name(),
	}, nil
}
