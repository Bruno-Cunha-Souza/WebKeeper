package main

import (
	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/database"
	_ "github.com/Bruno-Cunha-Souza/WebKeeper/internal/routes"
	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/services"
)

func main() {
	database.ConectDB()
	// routes.HandleRequests()
	services.StartMonit()
}
