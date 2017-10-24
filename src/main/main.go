package main

import (
	"net/http"

	. "config"
	. "database"
)

var (
	err error
)

func main() {
	Config, err = LoadConfiguration("./config.json")
	panicOnErr(err)
	Pool, err = InitDB(Config.Database.Provider, Config.Database.Host)
	panicOnErr(err)
	http.ListenAndServe(Config.Host+":"+Config.Port, nil)
	defer Pool.Close()
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
