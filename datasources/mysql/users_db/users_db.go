package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(fmt.Sprintf("Error loading .env file %s", err))
	}

	return os.Getenv(key)
}

var (
	Client *sql.DB

	username = goDotEnvVariable("MYSQL_USERS_USERNAME")
	password = goDotEnvVariable("MYSQL_USERS_PASSWORD")
	host     = goDotEnvVariable("MYSQL_USERS_HOST")
	port     = goDotEnvVariable("MYSQL_USERS_PORT")
	schema   = goDotEnvVariable("MYSQL_USERS_SCHEMA")
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		username, password, host, port, schema,
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
