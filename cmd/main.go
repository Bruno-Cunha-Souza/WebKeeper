package main

import (
	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/database"
	"github.com/Bruno-Cunha-Souza/WebKeeper/internal/services"
)

func main() {
	database.ConectDB()
	services.StartMonit()
}
