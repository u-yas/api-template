package settings

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// loading .env file and
// return remote or local development environment
func LoadEnv() *string {

	remote, ok := os.LookupEnv("REMOTE")

	if !ok {
		// Load dotenv from .env.local
		err := godotenv.Load(".env.local")
		if err != nil {
			panic("do not exists \".env.local\" in root folder. ")
		}
		fmt.Println("connect local to development")
	} else {
		if remote == "local" {
			err := godotenv.Load(".env.production.local")
			if err != nil {
				panic("do not exists \".env.production.local\" in root folder. ")
			}
			fmt.Println("connect remote to production")
		}
	}
	return &remote
}
