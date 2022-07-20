package main

import (
	"github.com/Rock2k3/notes-users/config"
	"github.com/Rock2k3/notes-users/internal/adapters"
	"github.com/Rock2k3/notes-users/internal/server"
	"log"
)

func main() {
	appConf, err := config.NewAppConfig().Load()
	if err != nil {
		log.Fatal(err)
	}

	err = adapters.NewPostgresDb(appConf).Init()
	if err != nil {
		log.Fatal(err)
	}

	server.NewServer(appConf).Run()
}
