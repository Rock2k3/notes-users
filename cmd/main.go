package main

import (
	"log"
	"notes-users/config"
	"notes-users/internal/adapters"
	"notes-users/internal/server"
)

func main() {
	appConf, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = adapters.NewPostgresDb(appConf).Init()
	if err != nil {
		log.Fatal(err)
	}

	server.NewServer(appConf).Run()
}
