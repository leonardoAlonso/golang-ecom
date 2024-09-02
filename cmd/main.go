package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"

	"github.com/leonardoAlonso/go-ecom/cmd/api"
	"github.com/leonardoAlonso/go-ecom/config"
	"github.com/leonardoAlonso/go-ecom/db"
)

func main() {
	db, err := db.NewMySqlStorage(
		mysql.Config{
			User:                 config.Envs.DBUser,
			Passwd:               config.Envs.DBPassword,
			Addr:                 config.Envs.DBAddress,
			DBName:               config.Envs.DBName,
			Net:                  "tcp",
			AllowNativePasswords: true,
			ParseTime:            true,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	InitStorage(db) // Pointer to the database connection

	server := api.NewApiServer(":8080", db) // Pointer to the struct
	if err := server.Run(); err != nil {    // Run on the same memory address
		log.Fatal(err)
	}
}

func InitStorage(db *sql.DB) {
	err := db.Ping() // check if the connection is alive
	if err != nil {
		log.Fatal(err)
	}

	// start the connection
	log.Println("Connected to the database")
}
