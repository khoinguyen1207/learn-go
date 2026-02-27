package mail

import (
	"context"
	"fmt"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
)

type SESProviderConfig struct {
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	From            string
}

type SESProvider struct {
	client *sesv2.Client
	cfg    *SESProviderConfig
}

func NewSESProvider(cfg *SESProviderConfig) (*SESProvider, error) {
	awsCfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithRegion(cfg.Region),
		awsConfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.AccessKeyID,
				cfg.SecretAccessKey,
				"",
			),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("ses: failed to load AWS config: %w", err)
	}

	return &SESProvider{
		client: sesv2.NewFromConfig(awsCfg),
		cfg:    cfg,
	}, nil
}

func (p *SESProvider) Name() string {
	return "ses"
}

func (p *SESProvider) SendMail(ctx context.Context, msg *MailMessage) (*SendResult, error) {
	msg.From = Address{
		Name:  "Support Team",
		Email: p.cfg.From,
	}
	return &SendResult{
		MessageID: "ses-message-id-placeholder",
		Provider:  p.Name(),
	}, nil
}
