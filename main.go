package main

import (
	"log"

	"github.com/marius004/url-shortener/database"
	"github.com/marius004/url-shortener/internal"
)

func main() {
	config := internal.NewConfig()
	db, err := database.ConnectToPSQL(database.GenerateDatabaseDSN(config))

	if err != nil {
		log.Fatalln(err)
	}

	server := NewServer(config, db, internal.NewLogger("logs.log"))
	server.Serve()
}
