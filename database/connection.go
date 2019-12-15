package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"monster_duel_api/utils"
	
	// Will automaticaly be used on gorm connection opening
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var connection *gorm.DB

type dbEnv struct {
	DBHost		string
	DBUser		string
	DBName		string
	DBPort		string
	DBPassword	string
}

func buildDBEnv() dbEnv {
	return dbEnv {
		DBHost: utils.GetEnvOrDefault("DB_HOST", "localhost"),
		DBPort: utils.GetEnvOrDefault("DB_PORT", "54322"),
		DBUser: utils.GetEnvOrDefault("DB_USER", "postgres"),
		DBName: utils.GetEnvOrDefault("DB_NAME", "yugi-database"),
		DBPassword: utils.GetEnvOrDefault("DB_PASSWORD", "postgres"),
	}
}

func openConnection() *gorm.DB {
	env := buildDBEnv()
	psqlEnv := fmt.Sprintf(`
		host=%s
		port=%s
		user=%s
		dbname=%s
		password=%s
		sslmode=disable
	`, env.DBHost, env.DBPort, env.DBUser, env.DBName, env.DBPassword)
	connection, err := gorm.Open("postgres",  psqlEnv)

	if err != nil {
		fmt.Println(err)
		panic("Could not connect to database.")
	}

	connection.LogMode(true)

	return connection
}

// GetDBConnection to get DB connection
func GetDBConnection() *gorm.DB {
	if connection == nil {
		connection = openConnection()
	}

	return connection
}
