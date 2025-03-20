package config

import (
	"database/sql"
	"employee-auth/models"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the MySQL database connection
func ConnectDatabase() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}

	// Read database credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Construct DSN (without database) to check for existence
	rootDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", dbUser, dbPass, dbHost, dbPort)
	dbDsn := fmt.Sprintf("%s%s?charset=utf8mb4&parseTime=True&loc=Local", rootDsn, dbName)

	// Connect using database/sql to check and create the database
	sqlDB, err := sql.Open("mysql", rootDsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL server: %v", err)
	}
	defer sqlDB.Close()

	// Check if database exists
	_, err = sqlDB.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	log.Println("✅ Database checked/created successfully")

	// Now, connect using GORM with the actual database
	DB, err = gorm.Open(mysql.Open(dbDsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Auto migrate models
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("✅ MySQL database connection successfully established")
}
