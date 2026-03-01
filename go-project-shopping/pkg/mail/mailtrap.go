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
type MailtrapResponse struct {
	Success    bool     `json:"success"`
	MessageIDs []string `json:"message_ids"`
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
	from := Address{
		Email: p.cfg.MailSender,
		Name:  p.cfg.NameSender,
	}
	content := struct {
		From    Address   `json:"from"`
		To      []Address `json:"to"`
		Cc      []Address `json:"cc,omitempty"`
		Subject string    `json:"subject"`
		Text    string    `json:"text,omitempty"`
		HTML    string    `json:"html,omitempty"`
	}{
		From:    from,
		To:      msg.To,
		Cc:      msg.Cc,
		Subject: msg.Subject,
		Text:    msg.BodyText,
		HTML:    msg.BodyHTML,
	}

	payload, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.cfg.BaseURL, bytes.NewReader(payload))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+p.cfg.APIKey)
	req.Header.Add("Content-Type", "application/json")

	// Send request
	res, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Read and parse response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read Mailtrap response body: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Mailtrap API error: %s", string(body))
	}

	var mailtrapResp MailtrapResponse
	if err = json.Unmarshal(body, &mailtrapResp); err != nil {
		return nil, fmt.Errorf("Failed to parse Mailtrap response: %v", err)
	}

	return &SendResult{
		MessageID: mailtrapResp.MessageIDs[0],
		Provider:  p.Name(),
	}, nil
}
