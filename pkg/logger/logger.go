package mblogger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type LogLevel string

const (
	InfoLevel    LogLevel = "info"
	ErrorLevel   LogLevel = "error"
	DebugLevel   LogLevel = "debug"
	WarningLevel LogLevel = "warning"
	FatalLevel   LogLevel = "fatal"
)

// LoggerParams contains the parameters for the logger service
type LoggerParams struct {
	UserID     string
	BotID      string
	SchemaName string
	TableName  string
}

// LoggerService is the service for logging messages to a database
type LoggerService struct {
	params        LoggerParams
	db            *gorm.DB
	consoleLogger *log.Logger
}

// LogRecord is the struct for the log record
type LogRecord struct {
	ID        uint `gorm:"primarykey"`
	Timestamp time.Time
	UserID    string         `gorm:"index:idx_user_bot_level,priority:1"`
	BotID     string         `gorm:"index:idx_user_bot_level,priority:2"`
	Level     LogLevel       `gorm:"index:idx_user_bot_level,priority:3"`
	Message   string         `gorm:"type:text"`
	Meta      datatypes.JSON `gorm:"type:jsonb"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
}

// NewLoggerService creates a new logger service
func NewLoggerService(params LoggerParams, db *gorm.DB) (*LoggerService, error) {
	service := &LoggerService{
		params:        params,
		db:            db,
		consoleLogger: log.New(os.Stdout, "", log.LstdFlags),
	}
	return service, service.autoMigrate()
}

// autoMigrate auto-migrates the table if it doesn't exist
func (s *LoggerService) autoMigrate() error {
	if s.db.Table(s.getTableName()).Migrator().HasTable(&LogRecord{}) {
		return nil
	}
	return s.db.Table(s.getTableName()).AutoMigrate(&LogRecord{})
}

// getTableName returns the fully qualified table name
func (s *LoggerService) getTableName() string {
	return fmt.Sprintf("%s.%s", s.params.SchemaName, s.params.TableName)
}

// log logs a message to the database and console
func (s *LoggerService) log(level LogLevel, message string, meta map[string]interface{}) error {
	s.logToConsole(level, message, meta)

	metaJSON, err := json.Marshal(meta)
	if err != nil {
		return fmt.Errorf("failed to marshal meta: %w", err)
	}

	record := LogRecord{
		Timestamp: time.Now(),
		UserID:    s.params.UserID,
		BotID:     s.params.BotID,
		Level:     level,
		Message:   message,
		Meta:      metaJSON,
	}

	if err := s.db.Table(s.getTableName()).Create(&record).Error; err != nil {
		return fmt.Errorf("failed to insert log into database: %w", err)
	}

	return nil
}

// logToConsole logs a message to the console
func (s *LoggerService) logToConsole(level LogLevel, message string, meta map[string]interface{}) {
	s.consoleLogger.Printf("[%s] %s %v", level, message, meta)
}

// Close closes the database connection
func (s *LoggerService) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}
	return sqlDB.Close()
}

// GetLogs retrieves logs from the database
func (s *LoggerService) GetLogs(limit int) ([]LogRecord, error) {
	var logs []LogRecord
	err := s.db.Table(s.getTableName()).
		Where("user_id = ? AND bot_id = ?", s.params.UserID, s.params.BotID).
		Order("created_at desc").
		Limit(limit).
		Find(&logs).Error
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve logs: %w", err)
	}
	return logs, nil
}

// GetLogsByLevel retrieves logs from the database by level
func (s *LoggerService) GetLogsByLevel(level LogLevel, limit int) ([]LogRecord, error) {
	var logs []LogRecord
	err := s.db.Table(s.getTableName()).
		Where("user_id = ? AND bot_id = ? AND level = ?", s.params.UserID, s.params.BotID, level).
		Order("created_at desc").
		Limit(limit).
		Find(&logs).Error
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve logs by level: %w", err)
	}
	return logs, nil
}

// Log methods (Info, Error, Debug, Warning, Fatal)
func (s *LoggerService) Info(message string, meta map[string]interface{}) {
	if err := s.log(InfoLevel, message, meta); err != nil {
		s.consoleLogger.Printf("Failed to log info message: %v", err)
	}
}

func (s *LoggerService) Error(message string, meta map[string]interface{}) {
	if err := s.log(ErrorLevel, message, meta); err != nil {
		s.consoleLogger.Printf("Failed to log error message: %v", err)
	}
}

func (s *LoggerService) Debug(message string, meta map[string]interface{}) {
	if err := s.log(DebugLevel, message, meta); err != nil {
		s.consoleLogger.Printf("Failed to log debug message: %v", err)
	}
}

func (s *LoggerService) Warning(message string, meta map[string]interface{}) {
	if err := s.log(WarningLevel, message, meta); err != nil {
		s.consoleLogger.Printf("Failed to log warning message: %v", err)
	}
}

func (s *LoggerService) Fatal(message string, meta map[string]interface{}) {
	if err := s.log(FatalLevel, message, meta); err != nil {
		s.consoleLogger.Printf("Failed to log fatal message: %v", err)
	}
}
