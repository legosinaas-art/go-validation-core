package configs

import (
	"bufio"
	"log"
	"os"
	"strings"
	"sync"
)

type Config struct {
	Port         string
	JWTSecret    []byte
	TemplatesDir string
}

var envOnce sync.Once

// loadEnvFile loads variables from a local .env file if it exists.
// Format: KEY=VALUE, lines starting with # are comments.
func loadEnvFile() {
	envOnce.Do(func() {
		f, err := os.Open(".env")
		if err != nil {
			return
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}

			parts := strings.SplitN(line, "=", 2)
			if len(parts) != 2 {
				continue
			}

			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])

			// Do not overwrite explicitly set env vars.
			if _, exists := os.LookupEnv(key); !exists && key != "" {
				_ = os.Setenv(key, val)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Printf("error reading .env: %v", err)
		}
	})
}

// Load reads configuration from environment variables and .env (if present).
func Load() *Config {
	loadEnvFile()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Println("JWT_SECRET is not set; using insecure default value")
		secret = "secretKey"
	}

	templatesDir := os.Getenv("TEMPLATES_DIR")
	if templatesDir == "" {
		templatesDir = "templates"
	}

	return &Config{
		Port:         port,
		JWTSecret:    []byte(secret),
		TemplatesDir: templatesDir,
	}
}



