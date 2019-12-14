package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	
	// Will automaticaly be used on gorm connection opening
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var connection *gorm.DB

func openConnection() *gorm.DB {
	// POSTGRES_USER: 'postgres'
	// POSTGRES_PASSWORD: 'postgres'
	// POSTGRES_DB: 'yugi-database'

	connection, err := gorm.Open("postgres", "host=localhost port=54322 user=postgres dbname=yugi-database password=postgres sslmode=disable")

	if err != nil {
		fmt.Println(err)
		panic("Could not connect to database.")
	}

	return connection
}

// GetDBConnection to get DB connection
func GetDBConnection() *gorm.DB {
	if connection == nil {
		connection = openConnection()
	}

	return connection
}
