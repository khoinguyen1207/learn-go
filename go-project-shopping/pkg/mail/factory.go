package mail

import (
	"fmt"
	"project-shopping/internal/config"
	"time"
)

type ProviderType string

const (
	ProviderSES      ProviderType = "ses"
	ProviderMailtrap ProviderType = "mailtrap"
)

func NewMailProvider(cfg *config.Config) (MailProvider, error) {
	switch ProviderType(cfg.MailProviderType) {
	case ProviderSES:
		return NewSESProvider(&SESProviderConfig{
			From:            cfg.Mail.FromAddress,
			Region:          cfg.AWS.Region,
			AccessKeyID:     cfg.AWS.AccessKeyID,
			SecretAccessKey: cfg.AWS.SecretAccessKey,
		})
	case ProviderMailtrap:
		return NewMailtrapProvider(&MailtrapProviderConfig{
			MailSender: cfg.Mail.FromAddress,
			NameSender: cfg.Mailtrap.NameSender,
			APIKey:     cfg.Mailtrap.APIKey,
			BaseURL:    cfg.Mailtrap.BaseURL,
			Timeout:    10 * time.Second,
		})
	default:
		return nil, fmt.Errorf("Unsupported mail provider: %s", cfg.MailProviderType)
	}
}
