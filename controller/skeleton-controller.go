package controller

import (
	"github.com/gin-gonic/gin"
	"go-restservice-skeleton/dto"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong! " + time.Now().Format(time.DateTime),
	})
}

func CreateSkeletonData(c *gin.Context) {
	var newSkeletonData dto.SkeletonData

	if err := c.ShouldBindJSON(&newSkeletonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Create SkeletonData with ID: " + strconv.Itoa(newSkeletonData.SID))
	c.JSON(http.StatusCreated, newSkeletonData)
}

func GetSkeletonData(c *gin.Context) {
	var skeletonData dto.SkeletonData
	skeletonData.SID = getSidFromRequest(c)
	skeletonData.Description = "SkeletonData with ID: " + strconv.Itoa(skeletonData.SID)
	log.Println("Get SkeletonData with ID: " + strconv.Itoa(skeletonData.SID))
	c.JSON(http.StatusOK, skeletonData)
}

func UpdateSkeletonData(c *gin.Context) {

	var updateSkeletonData dto.SkeletonData

	if err := c.ShouldBindJSON(&updateSkeletonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newSkeletonData dto.SkeletonData
	newSkeletonData.SID = getSidFromRequest(c)
	newSkeletonData.Description = updateSkeletonData.Description
	log.Println("Update SkeletonData with ID: " + strconv.Itoa(newSkeletonData.SID))
	c.JSON(http.StatusOK, newSkeletonData)
}

func DeleteSkeletonData(c *gin.Context) {
	skeletonDataID := getSidFromRequest(c)
	log.Println("Delete SkeletonData with ID: " + strconv.Itoa(skeletonDataID))
	c.Writer.WriteHeader(http.StatusNoContent)
}

func getSidFromRequest(c *gin.Context) int {
	id := c.Param("id")
	sid, _ := strconv.Atoi(id)
	return sid
}
