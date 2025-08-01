package main

import (
	"database/sql"
	"log"

	"github.com/ErZet99/SimpleBank/api"
	db "github.com/ErZet99/SimpleBank/db/sqlc"
	"github.com/ErZet99/SimpleBank/util"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
