package main

import (
	"github.com/afrizaloky/boilerplate/database"
	"github.com/afrizaloky/boilerplate/route"
)

func main() {

	db, _ := database.DBConnection()
	route.SetupRoutes(db)
}
