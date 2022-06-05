package main

import (
	"./db"
	"./route"
)

func main() {
	// Echo instance
	e := route.Init()
	db.Init()

	// Routes
	e.Logger.Fatal(e.Start(":3000"))
}
