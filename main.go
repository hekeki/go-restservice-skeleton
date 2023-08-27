package main

import (
	"github.com/gorilla/mux"
	"go-restservice-skeleton/controller"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/skeletonData", controller.CreateSkeletonData).Methods("POST")
	router.HandleFunc("/api/v1/skeletonData/{sid:[0-9]+}", controller.GetSkeletonData).Methods("GET")
	router.HandleFunc("/api/v1/skeletonData/{sid:[0-9]+}", controller.UpdateSkeletonData).Methods("PUT")
	router.HandleFunc("/api/v1/skeletonData/{sid:[0-9]+}", controller.DeleteSkeletonData).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8180", router))
}
