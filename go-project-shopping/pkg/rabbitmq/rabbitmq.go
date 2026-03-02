package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

type rabbitMQService struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	logger  *zerolog.Logger
}

func NewRabbitMQService(url string, logger *zerolog.Logger) RabbitMQService {
	conn, err := amqp091.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	return &rabbitMQService{
		conn:    conn,
		channel: ch,
		logger:  logger,
	}
}

func (r *rabbitMQService) Publish(ctx context.Context, queue string, message any) error {
	_, err := r.channel.QueueDeclare(
		queue, // queue name
		true,  // durable -> If true, the queue will survive a broker restart
		false, // delete when unused -> If true, the queue will be deleted when there are no more consumers
		false, // exclusive -> If true, the queue can only be used by the current connection and will be deleted when the connection closes
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("Failed to declare queue: %w", err)
	}

	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("Failed to marshal message: %w", err)
	}

	err = r.channel.PublishWithContext(ctx,
		"",    // exchange
		queue, // routing key
		false, // mandatory
		false, // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return fmt.Errorf("Failed to publish message: %w", err)
	}

	return nil
}

func (r *rabbitMQService) Consume(ctx context.Context, queue string, handler func([]byte) error) error {
	_, err := r.channel.QueueDeclare(
		queue, // queue name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("Failed to declare queue: %w", err)
	}

	msgs, err := r.channel.Consume(
		queue, // queue
		"",    // consumer
		false, // auto-ack -> If true, the server will consider messages acknowledged once delivered. If false, the server expects explicit acknowledgments from the consumer.
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return fmt.Errorf("Failed to register a consumer: %w", err)
	}

	go func() {
		for {
			select {
			case d, ok := <-msgs:
				if !ok {
					r.logger.Info().Msg("RabbitMQ channel closed")
					return
				}
				if err := handler(d.Body); err != nil {
					r.logger.Error().Err(err).Msg("Failed to handle message")
					d.Nack(false, false) // requeue the message
				} else {
					d.Ack(false)
				}
			case <-ctx.Done():
				r.logger.Info().Msg("Stopping RabbitMQ consumer")
				return
			}
		}
	}()

	return nil
}

func (r *rabbitMQService) Close() error {
	if r.channel != nil {
		if err := r.channel.Close(); err != nil {
			r.logger.Error().Err(err).Msg("Failed to close RabbitMQ channel")
		}
	}

	if r.conn != nil {
		if err := r.conn.Close(); err != nil {
			r.logger.Error().Err(err).Msg("Failed to close RabbitMQ connection")
		}
	}

	return nil
}
