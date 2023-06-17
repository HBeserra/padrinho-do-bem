package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

var DBConnection *sql.DB

func getMysqlConnectionUrl() string {
	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DB")

	var authString = fmt.Sprintf("%s:%s@", user, pass)
	if len(user) <= 0 {
		authString = ""
	}

	var connectionUrl = fmt.Sprintf("%s(%s:%s)/%s", authString, host, port, dbName)

	log.Println(connectionUrl)

	return connectionUrl
}

func Connect() {
	var connectionErr error
	var url = getMysqlConnectionUrl()

	DBConnection, connectionErr = sql.Open("mysql", url)
	if connectionErr != nil {
		log.Fatal(connectionErr)
	}

	err := DBConnection.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Define o timeout e limita o numero de conexões ao banco de dados
	DBConnection.SetConnMaxLifetime(time.Duration(10) * time.Second)
	DBConnection.SetMaxIdleConns(5)
	DBConnection.SetMaxOpenConns(2)
}
