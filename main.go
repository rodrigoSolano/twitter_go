package main

import (
	"log"

	"github.com/rodrigoSolano/twitter_go/db"
	"github.com/rodrigoSolano/twitter_go/handlers"
)

func main() {
	if !db.CheckConnection() {
		log.Fatal("Connection to database failed")
		return
	}
	handlers.Handlers()
}
