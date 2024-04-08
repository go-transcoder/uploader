package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"sync"
)

type Database struct {
	Host       string
	Port       string
	Name       string
	User       string
	Password   string
	SSLMode    string
	Connection *sql.DB
}

var lock = &sync.Mutex{}

var dbInstance *Database

func GetInstance() *Database {
	if dbInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if dbInstance == nil {
			fmt.Println("Creating db instance now.")
			dbInstance = &Database{
				Host:     os.Getenv("DBHOST"),
				Port:     os.Getenv("DBPORT"),
				Name:     os.Getenv("DBNAME"),
				User:     os.Getenv("DBUSER"),
				Password: os.Getenv("DBPASS"),
				SSLMode:  os.Getenv("SSLMODE"),
			}
			dbInstance.Connection, _ = dbInstance.getConnection()
		}
	}
	return dbInstance
}

func (_db Database) getConnection() (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", _db.User, _db.Password, _db.Host, _db.Port, _db.Name, _db.SSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
