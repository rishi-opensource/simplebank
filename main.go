package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/rishikant42/simplebank/api"
	db "github.com/rishikant42/simplebank/db/sqlc"
	"github.com/rishikant42/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Unable to load config")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to DB", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("unable to create new server", err.Error())
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err.Error())
	}
}
