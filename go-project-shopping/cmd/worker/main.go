package main

import (
	"context"
	"encoding/json"
	"log"
	"os/signal"
	"path/filepath"
	"project-shopping/internal/config"
	"project-shopping/internal/utils"
	"project-shopping/pkg/logger"
	"project-shopping/pkg/mail"
	"project-shopping/pkg/rabbitmq"
	"project-shopping/pkg/template"
	"sync"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

type Worker struct {
	rabbitMQ    rabbitmq.RabbitMQService
	mailService mail.MailService
	cfg         *config.Config
	logger      *zerolog.Logger
}

func NewWorker(cfg *config.Config) *Worker {
	// Initialize RabbitMQ service
	workerLogger := utils.NewLoggerWithPath("worker.log", "info")
	rabbitMQService, err := rabbitmq.NewRabbitMQService(cfg.RabbitMQURL, workerLogger)
	if err != nil {
		workerLogger.Fatal().Err(err).Msg("Failed to initialize RabbitMQ service")
	}

	// Initialize template service
	rootDir := utils.GetRootDir()
	templateDir := filepath.Join(rootDir, "pkg", "template")
	templateService := template.NewTemplateService(templateDir)

	// Initialize mail service
	mailProvider, err := mail.NewMailProvider(cfg)
	if err != nil {
		workerLogger.Fatal().Err(err).Msg("Failed to initialize mail provider")
	}
	mailLogger := utils.NewLoggerWithPath("mail.log", "info")
	mailService := mail.NewMailService(cfg, mailProvider, mailLogger, templateService)

	return &Worker{
		rabbitMQ:    rabbitMQService,
		mailService: mailService,
		cfg:         cfg,
		logger:      workerLogger,
	}
}

func (w *Worker) Start(ctx context.Context) error {
	log.Println("✅ Worker started, waiting for messages...")

	handlers := func(body []byte) error {
		w.logger.Info().Msgf("Received message: %s", string(body))

		var msg mail.MailMessageTemplate
		if err := json.Unmarshal(body, &msg); err != nil {
			w.logger.Error().Err(err).Msg("Failed to unmarshal message")
			return err
		}

		if err := w.mailService.SendWithTemplate(ctx, &msg); err != nil {
			w.logger.Error().Err(err).Msg("Failed to send email")
			return err
		}

		w.logger.Info().Msg("Email sent successfully")
		return nil
	}

	if err := w.rabbitMQ.Consume(ctx, rabbitmq.AuthEmailQueue, handlers); err != nil {
		w.logger.Fatal().Err(err).Msg("Failed to start consuming messages")
		return err
	}

	w.logger.Info().Msgf("Worker is running... Comsuming messages from queue: %s", rabbitmq.AuthEmailQueue)

	<-ctx.Done() // Wait for shutdown signal
	w.logger.Info().Msg("Worker is shutting down...")

	return ctx.Err()

}

func (w *Worker) Shutdown(ctx context.Context) error {
	w.logger.Info().Msg("Shutting down worker...")

	if err := w.rabbitMQ.Close(); err != nil {
		w.logger.Error().Err(err).Msg("Failed to close RabbitMQ connection")
		return err
	}

	// Handle all message before shutdown
	select {
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			w.logger.Warn().Msg("Worker shutdown timed out, some messages may not be processed")
		}
	default:
	}

	w.logger.Info().Msg("Worker shutdown complete")

	return nil
}

func main() {
	// Load environment variables from .env file
	rootDir := utils.GetRootDir()
	filepath := filepath.Join(rootDir, ".env")

	if err := godotenv.Load(filepath); err != nil {
		log.Fatalf("⚠️ Error loading .env file: %v", err)
	} else {
		log.Println("✅ .env file loaded successfully")
	}

	// Initialize configuration
	config.NewConfig()

	// Init logger
	logger.InitLogger(config.Get().AppEnv)

	// Initialize worker
	worker := NewWorker(config.Get())
	if worker == nil {
		log.Fatal("Failed to initialize worker")
	}

	// Start worker with graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := worker.Start(ctx); err != nil && err != context.Canceled {
			logger.Log.Fatal().Err(err).Msg("Worker encountered an error")
		}
	}()

	<-ctx.Done() // Wait for shutdown signal
	logger.Log.Info().Msg("⚠️ Shutdown signal received, stopping worker...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := worker.Shutdown(shutdownCtx); err != nil {
		logger.Log.Error().Err(err).Msg("❌ Worker shutdown failed")
	}

	wg.Wait()
	logger.Log.Info().Msg("✅ Worker stopped gracefully")
}
