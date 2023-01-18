package main

import (
	"net/http"
	"tea/api"
	"tea/config"
	"tea/database"
	"tea/domain"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	conf, _ := config.NewConfig("./config/config.yaml")
	repo, _ := database.NewMongoRepository(conf.Database.URL, conf.Database.DB, conf.Database.Timeout)
	service := domain.NewStoreService(repo)
	handler := api.NewStoreHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/store/{id}", handler.Get).Methods("GET")
	router.HandleFunc("/api/v1/store", handler.Post).Methods("POST")
	router.HandleFunc("/api/v1/store", handler.Put).Methods("PUT")
	router.HandleFunc("/api/v1/store/{id}", handler.Delete).Methods("DELETE")
	log.Info("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
