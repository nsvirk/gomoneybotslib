package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nsvirk/gomoneybotslib/internal/database"
	mbstate "github.com/nsvirk/gomoneybotslib/pkg/state"
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
		TableName: "state",
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

	// Initialize state service
	stateParams := mbstate.StateParams{
		UserID:     config.UserID,
		BotID:      config.BotID,
		SchemaName: config.Schema,
		TableName:  config.TableName,
	}

	stateService, err := mbstate.NewStateService(stateParams, db)
	if err != nil {
		log.Fatalf("Failed to create state service: %v", err)
	}
	defer func() {
		if err := stateService.Close(); err != nil {
			log.Printf("Error closing state service: %v", err)
		}
	}()

	// Example usage of the state service
	if err := exampleStateOperations(stateService); err != nil {
		log.Fatalf("Error during state operations: %v", err)
	}
}

func exampleStateOperations(state *mbstate.StateService) error {
	// Set initial state
	err := state.Set("key1", "value1", map[string]interface{}{"initial": true})
	if err != nil {
		return fmt.Errorf("failed to set initial state: %w", err)
	}
	fmt.Println("Initial state set")

	// Get state
	value, meta, err := state.Get("key1")
	if err != nil {
		return fmt.Errorf("failed to get state: %w", err)
	}
	fmt.Printf("Get result - Value: %s, Meta: %v\n", value, meta)

	// Update state
	err = state.Set("key1", "value2", map[string]interface{}{"updated": true})
	if err != nil {
		return fmt.Errorf("failed to update state: %w", err)
	}
	fmt.Println("State updated")

	// Get updated state
	value, meta, err = state.Get("key1")
	if err != nil {
		return fmt.Errorf("failed to get updated state: %w", err)
	}
	fmt.Printf("Get result after update - Value: %s, Meta: %v\n", value, meta)

	// Delete state
	err = state.Delete("key1")
	if err != nil {
		return fmt.Errorf("failed to delete state: %w", err)
	}
	fmt.Println("State deleted")

	// Try to get deleted state
	value, meta, err = state.Get("key1")
	if err != nil {
		fmt.Printf("Expected error after deletion: %v\n", err)
	} else {
		fmt.Printf("Unexpected: got value after deletion - Value: %s, Meta: %v\n", value, meta)
	}

	return nil
}
