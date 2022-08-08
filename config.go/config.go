package config

type DBConfig struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBTable    string
	DBPort     string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		DBUsername: "postgres",
		DBPassword: "sql1234",
		DBHost:     "localhost",
		DBTable:    "postgres",
		DBPort:     "5432",
	}
}
