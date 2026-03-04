package mail

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
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
	toAddresses := make([]string, len(msg.To))
	toCcAddresses := make([]string, len(msg.Cc))
	for i, addr := range msg.Cc {
		toCcAddresses[i] = addr.Email
	}

	for i, addr := range msg.To {
		toAddresses[i] = addr.Email
	}

	input := &sesv2.SendEmailInput{
		FromEmailAddress: aws.String(p.cfg.From),
		Destination: &types.Destination{
			ToAddresses: toAddresses,
			CcAddresses: toCcAddresses,
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Subject: &types.Content{
					Data:    aws.String(msg.Subject),
					Charset: aws.String("UTF-8"),
				},
				Body: buildBody(msg.BodyText, msg.BodyHTML),
			},
		},
	}

	result, err := p.client.SendEmail(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("SES SendEmail error: %w", err)
	}

	return &SendResult{
		MessageID: aws.ToString(result.MessageId),
		Provider:  p.Name(),
	}, nil
}

func buildBody(text, html string) *types.Body {
	body := &types.Body{}
	if text != "" {
		body.Text = &types.Content{
			Data:    aws.String(text),
			Charset: aws.String("UTF-8"),
		}
	}
	if html != "" {
		body.Html = &types.Content{
			Data:    aws.String(html),
			Charset: aws.String("UTF-8"),
		}
	}
	return body
}
