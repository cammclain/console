package config

import "github.com/joho/godotenv"

func Load() error {
	return godotenv.Load()
}

type Settings struct {
	TeamServerPort string `env:"TEAMSERVER_PORT"`
	JWTSecret      string `env:"JWT_SECRET"`

	// Postgres
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresPort     int    `env:"POSTGRES_PORT" envDefault:"5432"`
	PostgresUser     string `env:"POSTGRES_USER" envDefault:"user"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresDBName   string `env:"POSTGRES_DATABASE_NAME" envDefault:"c2"`
}
