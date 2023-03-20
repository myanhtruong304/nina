package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/nina/api/handler"
	"github.com/nina/api/routes"
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
	routes := routes.NewRoute(r, store, h)

	server := &Server{
		store:  store,
		router: &routes,
	}

	if err != nil {
		log.Fatal("can not start server:", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("can not start server", err)
	}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
