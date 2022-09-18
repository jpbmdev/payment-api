package config

import "os"

// -----------------------------------------------
// -- Struct to store enviroment variables
// -----------------------------------------------

type configSchema struct {
	DBConnectionString string
	DBName             string
	Port               string
}

var ConfigSchema configSchema

func LoadConfig() {
	ConfigSchema = configSchema{
		DBConnectionString: os.Getenv("DB_CONNECTION_STRING"),
		DBName:             os.Getenv("DB_NAME"),
		Port:               os.Getenv("PORT"),
	}
}
