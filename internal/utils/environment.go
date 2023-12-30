package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func LoadEnvironment() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error to get workdir:", err)
		log.Panic(err)
	}

	// Construir la ruta completa al archivo .env
	envFilePath := filepath.Join(currentDir+"/dollar-bot", ".env")

	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatal("Error to load file .env")
	}

	if err != nil {
		log.Panic(err)
	}
}
