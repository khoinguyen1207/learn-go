package rabbitmq

import "context"

type RabbitMQService interface {
	Publish(ctx context.Context, queue QueueName, message any) error
	Consume(ctx context.Context, queue QueueName, handler func([]byte) error) error
	Close() error
}
