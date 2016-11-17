package main

import (
	"database/sql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {

	db, err := sql.Open("postgres", "user=postgres dbname=crypto_mining_dashboard_dev password=password host=localhost sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	handler := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(":3001", handler))

}
