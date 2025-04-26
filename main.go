package main

import (
	"github.com/gin-gonic/gin"
	"go-restservice-skeleton/controller"
	"log"
)

func main() {

	router := gin.Default()
	router.Use(gin.Logger())
	router.GET("/api/v1/ping", controller.Ping)
	router.POST("/api/v1/skeletonData", controller.CreateSkeletonData)
	router.GET("/api/v1/skeletonData/:sid", controller.GetSkeletonData)
	router.PUT("/api/v1/skeletonData/:sid", controller.UpdateSkeletonData)
	router.DELETE("/api/v1/skeletonData/:sid", controller.DeleteSkeletonData)
	log.Fatal(router.Run(":8180"))
}
