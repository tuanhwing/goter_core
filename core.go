package gotercore

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadENV() string {
	//Get env name
	env := "develop.env"
	if env, ok := os.LookupEnv("ENV"); ok {
		fmt.Printf("ENV: %s\n", env)
	}

	//Load env file
	err := godotenv.Load(env)
	if err != nil {
		fmt.Printf("Error loading %s file", env)
	}

	return env
}
