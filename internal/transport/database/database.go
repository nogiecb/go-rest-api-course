package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// NewDatabase - returns a pointer to a new database connection
func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up new database connection...")

	dbUsername := goDotEnvVariable("DB_USERNAME")
	dbPassword := goDotEnvVariable("DB_PASSWORD")
	dbHost := goDotEnvVariable("DB_HOST")
	dbName := goDotEnvVariable("DB_NAME")
	dbPort := goDotEnvVariable("DB_PORT")

	connectString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbPort, dbName, dbUsername, dbPassword)
	fmt.Println(connectString)

	db, err := gorm.Open("postgres", connectString)
	if err != nil {
		fmt.Println("Error with DB connection")
		return db, err
	}

	if err := db.DB().Ping(); err != nil {
		return db, err
	}

	return db, nil
}
