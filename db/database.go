package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

type Db struct {
	*sql.DB
}

func New() (*Db, error) {

	err := godotenv.Load()

	if err != nil{
		fmt.Print("error getting values")
	}

	dbhost := os.Getenv("DB_HOST")
	dbport, err := strconv.Atoi(os.Getenv("DB_PORT"))
	dbuser := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	connString := ConnString(dbhost, dbport, dbuser, dbPassword,dbname)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	// check if connection is good
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

func ConnString(host string, port int, user string, password string ,dbName string) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
}
