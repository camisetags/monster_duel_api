package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var connection *gorm.DB

func openConnection() *gorm.DB {
	connection, err := gorm.Open("")

	if err != nil {
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
