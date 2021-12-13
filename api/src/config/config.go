package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// Port where the API will run
	Port = 0
	//DatabaseConnectionString for the database
	DatabaseConnectionString = ""
	// SecretKey is used the sign the token
	SecretKey []byte
)

//Load environment variables
func Load() {
	var error error
	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}
	Port, error = strconv.Atoi(os.Getenv("API_PORT"))
	if error != nil {
		Port = 9000
	}
	DatabaseConnectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
