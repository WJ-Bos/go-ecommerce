package main

import (
	"Ecom/cmd/api"
	"Ecom/config"
	"Ecom/db"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := db.NewMySqlStorage(mysql.Config{
		User:                 config.Envs.DbUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddr,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initDbConnection(db)

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initDbConnection(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Successfully connected to Database")
}
