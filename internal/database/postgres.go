package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectPostgres establishes a connection to the PostgreSQL database
func ConnectPostgres(dsn, schema, logLevel string) (*gorm.DB, error) {
	var level logger.LogLevel
	switch logLevel {
	case "silent":
		level = logger.Silent
	case "error":
		level = logger.Error
	case "warn":
		level = logger.Warn
	case "info":
		level = logger.Info
	default:
		level = logger.Info // Default to Info level
	}
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(level),
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Set connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection: %v", err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Create the schema if it doesn't exist
	createSchemaSql := fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schema)
	if err := db.Exec(createSchemaSql).Error; err != nil {
		panic("failed to create schema: " + err.Error())
	}

	return db, nil
}

// ClosePostgres closes the database connection
func ClosePostgres(db *gorm.DB) error {
	if db == nil {
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database connection: %v", err)
	}

	err = sqlDB.Close()
	if err != nil {
		return fmt.Errorf("failed to close database connection: %v", err)
	}

	return nil
}
