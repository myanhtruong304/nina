package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/nina/api"
	"github.com/nina/api/handler"
	db "github.com/nina/db/sqlc"
	"github.com/nina/util"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config", err)
	}

	fmt.Println(config)
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	h := handler.NewHandler(config)
	r := gin.Default()
	store := db.NewStore(conn)
	server := api.NewServer(r, store, h)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can not start server:", err)
	}

}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
