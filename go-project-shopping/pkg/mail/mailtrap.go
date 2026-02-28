package mail

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
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

func (p *MailtrapProvider) SendMail(ctx context.Context, msg *MailMessage) (SendResult, error) {
	msg.From = Address{
		Email: p.cfg.MailSender,
		Name:  p.cfg.NameSender,
	}

	payload, err := json.Marshal(msg)
	if err != nil {
		return SendResult{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.cfg.BaseURL, bytes.NewReader(payload))
	if err != nil {
		fmt.Println(err)
		return SendResult{}, err
	}
	req.Header.Add("Authorization", "Bearer "+p.cfg.APIKey)
	req.Header.Add("Content-Type", "application/json")

	res, err := p.client.Do(req)
	if err != nil {
		return SendResult{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return SendResult{}, fmt.Errorf("Mailtrap API error: %s", string(body))
	}

	return SendResult{
		MessageID: "mailtrap-message-id-placeholder",
		Provider:  p.Name(),
	}, nil
}
