package main

import (
	"github.com/gin-gonic/gin"
	//"github.com/gorilla/mux"
	gin2 "go-restservice-skeleton/controller/gin"
	//controller "go-restservice-skeleton/controller/mux"
	"log"
)

func main() {

	// gin router
	router := gin.Default()
	router.Use(gin.Logger())
	router.GET("/api/v1/ping", gin2.Ping)
	router.POST("/api/v1/skeletonData", gin2.CreateSkeletonData)
	router.GET("/api/v1/skeletonData/:sid", gin2.GetSkeletonData)
	router.PUT("/api/v1/skeletonData/:sid", gin2.UpdateSkeletonData)
	router.DELETE("/api/v1/skeletonData/:sid", gin2.DeleteSkeletonData)
	log.Fatal(router.Run(":8180"))

	// mux router
	//router := mux.NewRouter().StrictSlash(true)
	//router.HandleFunc("/api/v1/ping", controller.Ping).Methods("GET")
	//router.HandleFunc("/api/v1/skeletonData", controller.CreateSkeletonData).Methods("POST")
	//router.HandleFunc("/api/v1/skeletonData/{sid:[0-9]+}", controller.GetSkeletonData).Methods("GET")
	//router.HandleFunc("/api/v1/skeletonData/{sid:[0-9]+}", controller.UpdateSkeletonData).Methods("PUT")
	//router.HandleFunc("/api/v1/skeletonData/{sid:[0-9]+}", controller.DeleteSkeletonData).Methods("DELETE")
	//
	//log.Fatal(http.ListenAndServe(":8180", router))
}
