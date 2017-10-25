package main

import (
	"net/http"

	"config"
	"models"
)

func main() {
	conf, err := config.LoadConfiguration("./config.json")
	panicOnErr(err)

	db, err := models.InitDB(conf.Database.Provider, conf.Database.Host)
	panicOnErr(err)

	defer db.Close()

	http.ListenAndServe(conf.Host+":"+conf.Port, nil)
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
