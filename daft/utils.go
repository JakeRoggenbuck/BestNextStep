package main

import (
	"fmt"
	"log"
	"os"
)

func GetEnvOrDefault(env_var string, default_var string) string {
	item := os.Getenv(env_var)
	if item == "" {
		return default_var
	} else {
		return item
	}
}

func GetEnvOrFatal(env_var string) string {
	item := os.Getenv(env_var)
	if item == "" {
		fmt.Printf("%s not set", env_var)
		log.Fatalf("%s not set", env_var)
	}

	return item
}
