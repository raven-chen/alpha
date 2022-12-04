package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/raven-chen/alpha/admin"
	"github.com/raven-chen/alpha/database"
	"github.com/raven-chen/alpha/orders"
)

func main() {
	// Setup project
	db := database.InitDB()
	mux := http.NewServeMux()
	b := admin.Initialize(db, mux)
	orders.Init(db, mux, b)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	fmt.Println("Served at http://localhost:" + port + "/admin")

	http.Handle("/", mux)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
