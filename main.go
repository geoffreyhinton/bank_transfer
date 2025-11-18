package main

import (
	"database/sql"
	"log"

	"github.com/geoffreyhinton/bank_transfer/api"
	db "github.com/geoffreyhinton/bank_transfer/db/sqlc"
	"github.com/geoffreyhinton/bank_transfer/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
