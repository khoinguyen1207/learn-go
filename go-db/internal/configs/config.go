package configs

import "os"

type DatabaseConfig struct {
	ConnectionString string
}

type Config struct {
	Db DatabaseConfig
}

func NewConfig() *Config {
	return &Config{
		Db: DatabaseConfig{
			ConnectionString: GetEnv("DATABASE_URL", "postgres://user:password@localhost:5432/mydb"),
		},
	}
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
