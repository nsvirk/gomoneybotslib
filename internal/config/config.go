package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	KiteUserId        string `env:"KITE_USER_ID"`
	KitePassword      string `env:"KITE_PASSWORD"`
	KiteTotpSecret    string `env:"KITE_TOTP_SECRET"`
	PostgresDsn       string `env:"POSTGRES_DSN"`
	PostgresSchema    string `env:"POSTGRES_SCHEMA"`
	PostgresLogLevel  string `env:"POSTGRES_LOG_LEVEL"`
	TelegramBotToken  string `env:"TELEGRAM_BOT_TOKEN"`
	TelegramChatId    string `env:"TELEGRAM_CHAT_ID"`
	BotId             string `env:"BOT_ID"`
	BotDesc           string `env:"BOT_DESC"`
	BotLogsTablename  string `env:"BOT_LOGS_TABLENAME"`
	BotStateTablename string `env:"BOT_STATE_TABLENAME"`
	BotSegmentExpiry  string `env:"BOT_SEGMENT_EXPIRY"`
	BotOCExchange     string `env:"BOT_OC_EXCHANGE"`
	BotOCName         string `env:"BOT_OC_NAME"`
	BotOCFutExpiry    string `env:"BOT_OC_FUT_EXPIRY"`
	BotOCOptExpiry    string `env:"BOT_OC_OPT_EXPIRY"`
}

var SingleLine string = "--------------------------------------------------"

func init() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// LoadEnv loads the environment variables
func LoadEnv() (*Config, error) {
	cfg, err := loadEnv()
	if err != nil {
		return nil, err
	}
	cfg.checkEnv()
	return cfg, nil
}

// loadEnv sets the values from the environment variables, using env tags
func loadEnv() (*Config, error) {
	var cfg Config
	// Get the value of the struct
	configValue := reflect.ValueOf(&cfg).Elem()
	// Get the type of the struct
	configType := configValue.Type()
	// Set the values from the environment variables
	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		fieldValue := configValue.Field(i)
		if field.PkgPath == "" {
			envKey := field.Tag.Get("env")
			if envKey != "" {
				envValue := os.Getenv(envKey)
				if envValue != "" {
					fieldValue.SetString(envValue)
				}
			}
		}
	}
	return &cfg, nil
}

// Check if all the env variables are set using env tags in struct
func (c *Config) checkEnv() {
	env := reflect.ValueOf(c).Elem()
	for i := 0; i < env.NumField(); i++ {
		envKey := env.Type().Field(i).Tag.Get("env")
		if envKey == "" {
			continue
		}
		if env.Field(i).String() == "" {
			log.Fatalf("Environment variable %s is not set", envKey)
		}
	}
}

// String returns the configuration as a string
func (c *Config) String() string {
	var sb strings.Builder

	sb.WriteString(SingleLine + "\n")
	sb.WriteString("Config: \n")
	sb.WriteString(SingleLine + "\n")

	t := reflect.TypeOf(*c)
	v := reflect.ValueOf(*c)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).String()

		// Mask sensitive fields
		value = maskSensitiveField(field.Name, value)
		sb.WriteString(fmt.Sprintf("  %s:  %s\n", field.Name, value))
	}

	sb.WriteString(SingleLine + "\n")

	return sb.String()
}

func maskSensitiveField(fieldName, value string) string {
	sensitiveFields := []string{"token", "dsn", "secret", "password", "url"}

	fieldNameLower := strings.ToLower(fieldName)
	for _, sensitive := range sensitiveFields {
		if strings.Contains(fieldNameLower, sensitive) {
			return maskValue(value)
		}
	}

	return value
}

func maskValue(value string) string {
	if len(value) <= 3 {
		return strings.Repeat("*", 7)
	}
	return value[:3] + strings.Repeat("*", 7)
}
