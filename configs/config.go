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
	Verify Verify
}

type Server struct {
	Addr string
}

type DbConfig struct {
	Dns string
}

type SMTP struct {
	From string
}

type Verify struct {
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
			From: os.Getenv("EMAIL_FROM"),
		},

		Verify: Verify{
			Email:    os.Getenv("EMAIL_TO"),
			Password: os.Getenv("PASSWORD"),
			Address:  os.Getenv("ADDRESS"),
		},
	}
}
