package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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
	err := godotenv.Load(".env")

	if err != nil {
		log.Println(err.Error())
	}

	return &Config{
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
}
