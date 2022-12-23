package Config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func loadENV() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error Load Env")
		os.Exit(1)
	}
}

func GetEnv(key string) string {
	loadENV()
	return os.Getenv(key)
}
