package configs

import (
	"bufio"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	instance *Config
	only     sync.Once
)

type Config struct {
	Db     DbConfig
	Server Server
	SMTP   SMTP
}

type Server struct {
	Addr string
}

type DbConfig struct {
	Dns string
}

// Email - smtp.gmail.com:587
// Email - данные авторизации
// Password - данные авторизации
type SMTP struct {
	Email    string
	Password string
	Address  string
}

func LoadConfig() *Config {

	only.Do(func() {
		loadDotEnv(".env")

		instance = &Config{
			Db: DbConfig{
				Dns: os.Getenv("DNS"),
			},

			Server: Server{
				Addr: os.Getenv("Addr"),
			},

			SMTP: SMTP{
				Email:    os.Getenv("EMAIL_TO"),
				Password: os.Getenv("PASSWORD"),
				Address:  os.Getenv("ADDRESS"),
			},
		}
	})

	return instance
}

func loadDotEnv(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
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
		value := strings.TrimSpace(parts[1])
		value = strings.Trim(value, `"'`)
		if key == "" {
			continue
		}

		// Keep already provided environment variables unchanged.
		if _, exists := os.LookupEnv(key); !exists {
			_ = os.Setenv(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}
}
