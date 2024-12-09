package config

import (
	"os"

	"github.com/joho/godotenv"
)

var isLoaded bool = false

func _init() {
	if !isLoaded {
		isdev := os.Getenv("PROD")
		if isdev != "true" {
			godotenv.Load(".env")
		}
	}
}

func GetServiceDomain(service string) string {
	_init()
	switch service {
	case "user":
		if data := os.Getenv("USER_ENDPOINT"); data != "" {
			return data
		}
		return "http://localhost:3000"
	case "product":
		if data := os.Getenv("PRODUCT_ENDPOINT"); data != "" {
			return data
		}
		return "http://localhost:3001"
	case "sale":
		if data := os.Getenv("SALE_ENDPOINT"); data != "" {
			return data
		}
		return "http://localhost:3002"
	default:
		return ""
	}
}

func GetDBConfig() (string, string) {
	_init()
	if driver, url := os.Getenv("DB_NAME"), os.Getenv("DB_URL"); url != "" && driver != "" {
		return driver, url
	}
	return "mysql", "root:secret@/default"
}
