package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nsvirk/gomoneybotslib/internal/database"
	mblogger "github.com/nsvirk/gomoneybotslib/pkg/logger"
)

func main() {
	// Configuration
	config := struct {
		DSN       string
		Schema    string
		TableName string
		LogLevel  string
		UserID    string
		BotID     string
	}{
		DSN:       os.Getenv("POSTGRES_DSN"),
		Schema:    "bots",
		TableName: "logger",
		LogLevel:  "error",
		UserID:    "SA0123",
		BotID:     "BOTv1",
	}

	// Initialize Postgres connection
	db, err := database.ConnectPostgres(config.DSN, config.Schema, config.LogLevel)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		if err := database.ClosePostgres(db); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	// Initialize logger service
	loggerParams := mblogger.LoggerParams{
		UserID:     config.UserID,
		BotID:      config.BotID,
		SchemaName: config.Schema,
		TableName:  config.TableName,
	}

	logger, err := mblogger.NewLoggerService(loggerParams, db)
	if err != nil {
		log.Fatalf("Failed to create logger service: %v", err)
	}
	defer func() {
		if err := logger.Close(); err != nil {
			log.Printf("Error closing logger service: %v", err)
		}
	}()

	// Example usage of the logger service
	if err := exampleLoggerOperations(logger); err != nil {
		log.Fatalf("Error during logger operations: %v", err)
	}
}

func exampleLoggerOperations(logger *mblogger.LoggerService) error {
	// Log messages
	logger.Info("Hello, info!", nil)
	logger.Error("Hello, error!", map[string]interface{}{"error_type": "InputException"})
	logger.Debug("Hello, debug!", map[string]interface{}{"debug_type": "logger"})
	logger.Warning("Hello, warning!", map[string]interface{}{"warn_type": "simple"})
	logger.Fatal("Hello, fatal!", map[string]interface{}{"fatal_type": "simple"})

	// Get logs
	logs, err := logger.GetLogs(5)
	if err != nil {
		return fmt.Errorf("failed to get logs: %w", err)
	}
	fmt.Println("\n\nAll logs:")
	for _, log := range logs {
		fmt.Printf("%+v\n", log)
	}

	// Get logs by level
	infoLogs, err := logger.GetLogsByLevel(mblogger.InfoLevel, 1)
	if err != nil {
		return fmt.Errorf("failed to get info logs: %w", err)
	}
	fmt.Println("\n\nInfo logs:")
	for _, log := range infoLogs {
		fmt.Printf("%+v\n", log)
	}

	return nil
}
