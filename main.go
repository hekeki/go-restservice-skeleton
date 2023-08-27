package main

import (
	"github.com/gorilla/mux"
	"go-restservice-skeleton/data"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/skeletonData", data.CreateSkeletonData).Methods("POST")
	router.HandleFunc("/api/v1/skeletonData/{sid:[0-9]+}", data.GetSkeletonData).Methods("GET")
	router.HandleFunc("/api/v1/skeletonData/{sid:[0-9]+}", data.UpdateSkeletonData).Methods("PUT")
	router.HandleFunc("/api/v1/skeletonData/{sid:[0-9]+}", data.DeleteSkeletonData).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8180", router))
}
