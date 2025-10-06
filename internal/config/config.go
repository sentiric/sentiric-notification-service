// sentiric-notification-service/internal/config/config.go
package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	GRPCPort string
	HttpPort string
	CertPath string
	KeyPath  string
	CaPath   string
	LogLevel string
	Env      string

	// Bildirim servisi bağımlılıkları (Placeholder)
	SmsAdapter      string // Twilio, Vonage, etc.
	EmailAdapter    string // SendGrid, Mailgun, etc.
	TwilioAuthToken string
}

func Load() (*Config, error) {
	godotenv.Load()

	// Harmonik Mimari Portlar (Yatay Yetenek, 172XX bloğu atandı)
	return &Config{
		GRPCPort: GetEnv("NOTIFICATION_SERVICE_GRPC_PORT", "17211"),
		HttpPort: GetEnv("NOTIFICATION_SERVICE_HTTP_PORT", "17210"),

		CertPath: GetEnvOrFail("NOTIFICATION_SERVICE_CERT_PATH"),
		KeyPath:  GetEnvOrFail("NOTIFICATION_SERVICE_KEY_PATH"),
		CaPath:   GetEnvOrFail("GRPC_TLS_CA_PATH"),
		LogLevel: GetEnv("LOG_LEVEL", "info"),
		Env:      GetEnv("ENV", "production"),

		SmsAdapter:      GetEnv("SMS_ADAPTER", "twilio"),
		EmailAdapter:    GetEnv("EMAIL_ADAPTER", "sendgrid"),
		TwilioAuthToken: GetEnv("TWILIO_AUTH_TOKEN", ""),
	}, nil
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func GetEnvOrFail(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal().Str("variable", key).Msg("Gerekli ortam değişkeni tanımlı değil")
	}
	return value
}
