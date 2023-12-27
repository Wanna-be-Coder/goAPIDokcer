package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/teachschool/simplebank/api"
	db "github.com/teachschool/simplebank/db/sqlc"
)

const (
	dbDriver       = "postgres"
	dbSource       = "postgresql://root:Asdf1234!@localhost:5432/simple_bank?sslmode=disable"
	serverAdddress = "0.0.0.0:8000"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can not connect to db", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAdddress)
	if err != nil {
		log.Fatal("Can not start server", err)
	}
}
