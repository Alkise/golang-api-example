package main

import (
	"net/http"

	. "config"
	. "database"
)

func main() {
	Config = LoadConfiguration("./config.json")
	Pool = InitDB(Config.Database.Provider, Config.Database.Host)
	http.ListenAndServe(Config.Host+":"+Config.Port, nil)
	defer Pool.Close()
}
