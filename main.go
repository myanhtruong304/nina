package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/nina/api"
	db "github.com/nina/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:nina@localhost:5435/nina_app?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("can not start server:", err)
	}

}
