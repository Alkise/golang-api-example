package main

import (
	"database/sql"
	"net/http"

	. "config"
	"models"
)

func main() {
	var err error

	Config, err = LoadConfiguration("./config.json")
	panicOnErr(err)

	var db *sql.DB
	db, err = models.InitDB(Config.Database.Provider, Config.Database.Host)
	panicOnErr(err)

	defer db.Close()

	http.ListenAndServe(Config.Host+":"+Config.Port, nil)
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
