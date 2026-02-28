package mail

import "context"

type MailProvider interface {
	SendMail(ctx context.Context, msg *MailMessage) (SendResult, error)
	Name() string
}

type MailService interface {
	SendMail(ctx context.Context, msg *MailMessage) error
}
