package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wipdev-tech/fcc-go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)

    fmt.Println("Listening at port 8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
