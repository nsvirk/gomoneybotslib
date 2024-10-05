package mbstate

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// StateParams are the parameters for the state service
type StateParams struct {
	UserID     string
	BotID      string
	SchemaName string
	TableName  string
}

type StateService struct {
	params StateParams
	db     *gorm.DB
}

type StateRecord struct {
	ID        uint           `gorm:"primarykey"`
	UserID    string         `gorm:"index:idx_user_bot_key,unique,priority:1"`
	BotID     string         `gorm:"index:idx_user_bot_key,unique,priority:2"`
	Key       string         `gorm:"index:idx_user_bot_key,unique,priority:3"`
	Value     string         `gorm:"type:text"`
	Meta      datatypes.JSON `gorm:"type:jsonb"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
}

// NewStateService creates a new state service
func NewStateService(params StateParams, db *gorm.DB) (*StateService, error) {
	service := &StateService{
		params: params,
		db:     db,
	}
	return service, service.AutoMigrate()
}

// AutoMigrate creates or updates the state table schema
func (s *StateService) AutoMigrate() error {
	if s.db.Table(s.getTableName()).Migrator().HasTable(&StateRecord{}) {
		return nil
	}
	return s.db.Table(s.getTableName()).AutoMigrate(&StateRecord{})
}

// getTableName returns the fully qualified table name
func (s *StateService) getTableName() string {
	return fmt.Sprintf("%s.%s", s.params.SchemaName, s.params.TableName)
}

// Get retrieves the value and metadata for a given key
func (s *StateService) Get(key string) (string, map[string]interface{}, error) {
	var record StateRecord
	result := s.db.Table(s.getTableName()).Where("user_id = ? AND bot_id = ? AND key = ?", s.params.UserID, s.params.BotID, key).First(&record)
	if result.Error != nil {
		return "", nil, result.Error
	}

	var meta map[string]interface{}
	if err := json.Unmarshal(record.Meta, &meta); err != nil {
		return "", nil, fmt.Errorf("failed to unmarshal metadata: %w", err)
	}

	return record.Value, meta, nil
}

// Set upserts the value and metadata for a given key
func (s *StateService) Set(key, value string, meta map[string]interface{}) error {
	jsonbMeta, err := json.Marshal(meta)
	if err != nil {
		return fmt.Errorf("failed to marshal metadata: %w", err)
	}

	record := &StateRecord{
		UserID:    s.params.UserID,
		BotID:     s.params.BotID,
		Key:       key,
		Value:     value,
		Meta:      datatypes.JSON(jsonbMeta),
		UpdatedAt: time.Now(),
	}

	return s.db.Table(s.getTableName()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "bot_id"}, {Name: "key"}},
		DoUpdates: clause.AssignmentColumns([]string{"value", "meta", "updated_at"}),
	}).Create(record).Error
}

// Delete removes the record for a given key
func (s *StateService) Delete(key string) error {
	return s.db.Table(s.getTableName()).Where("user_id = ? AND bot_id = ? AND key = ?", s.params.UserID, s.params.BotID, key).Delete(&StateRecord{}).Error
}

// Close closes the database connection
func (s *StateService) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return fmt.Errorf("failed to get database instance: %w", err)
	}
	return sqlDB.Close()
}
