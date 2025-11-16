package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	// godotenv.Load() carga las variables de entorno desde un archivo .env en el directorio actual
	error := godotenv.Load()
	if error != nil {
		log.Println("Error loading .env file")
	}
}

func GetEnv(key string, defaultValue string) string {
	// os.Getenv obtiene el valor de la variable de entorno SERVER_PORT
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
