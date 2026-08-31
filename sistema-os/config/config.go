package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresHost     string
	PostgresPort     string

	MySQLUser     string
	MySQLPassword string
	MySQLDB       string
	MySQLHost     string
	MySQLPort     string
}

func Carregar() Config {
	err := godotenv.Load(".env")

	if err != nil {
		err = godotenv.Load("../.env")
	}

	if err != nil {
		panic("Erro ao carregar o arquivo .env")
	}

	return Config{
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDB:       os.Getenv("POSTGRES_DB"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),

		MySQLUser:     os.Getenv("MYSQL_USER"),
		MySQLPassword: os.Getenv("MYSQL_PASSWORD"),
		MySQLDB:       os.Getenv("MYSQL_DB"),
		MySQLHost:     os.Getenv("MYSQL_HOST"),
		MySQLPort:     os.Getenv("MYSQL_PORT"),
	}
}
