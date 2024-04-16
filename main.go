package main

import (
	"log"
	"net/http"
	"nsq-mvc-example/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// Register handlers
	r.HandleFunc("/produce", controllers.ProduceMessageHandler).Methods("POST")
	r.HandleFunc("/consume", controllers.ConsumeMessageHandler).Methods("GET")

	log.Println("Starting NSQ Example...")
	log.Fatal(http.ListenAndServe(":8088", r))
}
