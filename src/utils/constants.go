package utils

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Token               string
	AdminChatIDEvg      string
	AdminNameEvg        string
	AdminChatIDNat      string
	AdminNameNat        string
	RabbitMQHost        string
	CoachChatIDs        []string
	RabbitMQDefaultUser string
	RabbitMQDefaultPass string
	Version             string
	BackendAPI          string
	EmailPattern        *regexp.Regexp
	PhonePattern        *regexp.Regexp
	MessagesDir         string
	MessagesFile        string
	MessagesPath        string
	MaxMessageLength    int
}

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	return &Config{
		Token:               getEnv("TELEGRAM_TOKEN"),
		AdminChatIDEvg:      getEnv("ADMIN_CHAT_ID_EVG"),
		AdminNameEvg:        getEnv("ADMIN_NAME_EVG"),
		AdminChatIDNat:      getEnv("ADMIN_CHAT_ID_NAT"),
		AdminNameNat:        getEnv("ADMIN_NAME_NAT"),
		RabbitMQHost:        getEnv("RABBITMQ_HOST"),
		CoachChatIDs:        strings.Split(getEnv("COACH_CHAT_IDS"), ","),
		RabbitMQDefaultUser: getEnv("RABBITMQ_DEFAULT_USER"),
		RabbitMQDefaultPass: getEnv("RABBITMQ_DEFAULT_PASS"),
		Version:             getEnv("VERSION"),
		BackendAPI:          getEnv("BACKEND_API"),
		EmailPattern:        regexp.MustCompile(`[^@]+@[^@]+\.[^@]+`),
		PhonePattern:        regexp.MustCompile(`^(?:\+7|8)[- ]?(?:\d{3}[- ]?\d{3}[- ]?\d{2}[- ]?\d{2})$`),
		MessagesDir:         "messages",
		MessagesFile:        "messages.json",
		MessagesPath:        "messages/messages.json",
		MaxMessageLength:    4090,
	}
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s is not found in environment", key)
	}
	return value
}
