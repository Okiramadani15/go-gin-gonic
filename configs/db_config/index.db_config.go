package db_config

import "os"

var DB_DRIVER = "postgres"

var DB_HOST = "localhost" //http://localhost

var DB_PORT = "5432"

var DB_NAME = "postgres"

var DB_USER = "postgres"

var DB_PASSWORD = "learngo123"

func InitDatabaseConfig() {

	driverEnv := os.Getenv("DB_DRIVER")

	if driverEnv != "" {
		DB_DRIVER = driverEnv
	}

	hostEnv := os.Getenv("DB_HOST")

	if hostEnv != "" {
		DB_HOST = hostEnv
	}

	portEnv := os.Getenv("DB_PORT")

	if portEnv != "" {
		DB_PORT = portEnv
	}

	nameEnv := os.Getenv("DB_NAME")

	if nameEnv != "" {
		DB_NAME = nameEnv
	}

	userEnv := os.Getenv("DB_USER")

	if userEnv != "" {
		DB_USER = userEnv
	}

	passwordEnv := os.Getenv("DB_PASSWORD")

	if passwordEnv != "" {
		DB_PASSWORD = passwordEnv
	}

}
