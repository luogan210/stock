package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// AppConfig holds all application configuration loaded from environment variables
// with sane defaults. It supports loading from a .env file in development.
type AppConfig struct {
	Env          string
	HTTPPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	TrustProxy   bool

	// Auth
	JWTSecret        string
	JWTExpireMinutes int

	// Uploads
	UploadDir          string
	MaxUploadSizeBytes int64
	ChunkSizeBytes     int
	UploadSessionTTL   time.Duration

	// Wechat (re-export)
	Wechat WechatConfig
}

var cfg *AppConfig

// Load loads configuration from environment variables and .env file (if present).
// It is safe to call multiple times; subsequent calls return the same instance.
func Load() *AppConfig {
	if cfg != nil {
		return cfg
	}

	// Load .env files if present (best effort, multiple locations)
	_ = godotenv.Load(
		".env",
		".env.local",
		"config/.env",
		"config/app.env",
		"config/dev.env",
	)

	cfg = &AppConfig{
		Env:          getEnv("APP_ENV", "development"),
		HTTPPort:     getEnv("HTTP_PORT", "8080"),
		ReadTimeout:  getEnvDuration("HTTP_READ_TIMEOUT", 15*time.Second),
		WriteTimeout: getEnvDuration("HTTP_WRITE_TIMEOUT", 15*time.Second),
		IdleTimeout:  getEnvDuration("HTTP_IDLE_TIMEOUT", 60*time.Second),
		TrustProxy:   getEnvBool("TRUST_PROXY", false),

		JWTSecret:        getEnv("JWT_SECRET", "change_me_in_env"),
		JWTExpireMinutes: getEnvInt("JWT_EXPIRE_MINUTES", 60*24), // 1 day default

		UploadDir:          getEnv("UPLOAD_DIR", "uploads"),
		MaxUploadSizeBytes: getEnvInt64("MAX_UPLOAD_SIZE_BYTES", 5*1024*1024*1024), // 5GB
		ChunkSizeBytes:     getEnvInt("CHUNK_SIZE_BYTES", 2*1024*1024),             // 2MB
		UploadSessionTTL:   getEnvDuration("UPLOAD_SESSION_TTL", 24*time.Hour),

		Wechat: *GetWechatConfig(),
	}

	return cfg
}

// Helpers

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
		log.Printf("WARN: invalid int for %s=%q, using default %d", key, v, defaultValue)
	}
	return defaultValue
}

func getEnvInt64(key string, defaultValue int64) int64 {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return i
		}
		log.Printf("WARN: invalid int64 for %s=%q, using default %d", key, v, defaultValue)
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if v := os.Getenv(key); v != "" {
		switch v {
		case "1", "true", "TRUE", "True", "yes", "Y", "y":
			return true
		case "0", "false", "FALSE", "False", "no", "N", "n":
			return false
		default:
			log.Printf("WARN: invalid bool for %s=%q, using default %v", key, v, defaultValue)
		}
	}
	return defaultValue
}

func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		d, err := time.ParseDuration(v)
		if err == nil {
			return d
		}
		log.Printf("WARN: invalid duration for %s=%q, using default %s", key, v, defaultValue)
	}
	return defaultValue
}
