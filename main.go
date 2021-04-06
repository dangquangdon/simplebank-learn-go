package main

import (
	"database/sql"
	"log"

	"github.com/dangquangdon/simplebank/api"
	db "github.com/dangquangdon/simplebank/db/sqlc"
	"github.com/dangquangdon/simplebank/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig("./environments", "dev")
	if err != nil {
		log.Fatal("Cannot load configuration: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start the server: ", err)
	}
}
