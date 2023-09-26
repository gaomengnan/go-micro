package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gaomengnan/go-micro/api"
	db "github.com/gaomengnan/go-micro/db/sqlc"
	"github.com/gaomengnan/go-micro/util"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")

	if err != nil {
		fmt.Println("err:%v", err)
		log.Fatal(err)
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
