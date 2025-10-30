package main


import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/michaelNuel/Book-Store-Management-api/PKG/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}