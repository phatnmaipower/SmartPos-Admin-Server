package main

import (
	"app/db"
	"app/route"
)

func main() {
	// Echo instance
	e := route.Init()
	db.Init()

	// Routes
	e.Logger.Fatal(e.Start(":3000"))
}
