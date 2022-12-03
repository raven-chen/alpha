package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/raven-chen/alpha/admin"
)

func main() {
	// Setup project
	mux := admin.Initialize()

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
