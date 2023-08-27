package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-restservice-skeleton/dto"
	"io"
	"log"
	"net/http"
	"strconv"
)

func CreateSkeletonData(writer http.ResponseWriter, request *http.Request) {
	var newSkeletonData dto.SkeletonData
	reqBody, errReadAll := io.ReadAll(request.Body)
	if errReadAll != nil {
		fmt.Println("error:", errReadAll)
	}

	errUnmarshal := json.Unmarshal(reqBody, &newSkeletonData)
	if errUnmarshal != nil {
		fmt.Println("error:", errUnmarshal)
	}

	writer.WriteHeader(http.StatusCreated)
	log.Println("Create SkeletonData with ID: " + strconv.Itoa(newSkeletonData.SID))

	errEncode := json.NewEncoder(writer).Encode(newSkeletonData)
	if errEncode != nil {
		fmt.Println("error:", errEncode)
	}
}

func GetSkeletonData(writer http.ResponseWriter, request *http.Request) {
	skeletonDataID := getSidFromRequest(request)
	writer.WriteHeader(http.StatusOK)

	var newSkeletonData dto.SkeletonData
	newSkeletonData.SID = skeletonDataID
	newSkeletonData.Description = "SkeletonData with ID: " + strconv.Itoa(skeletonDataID)
	log.Println("Get SkeletonData with ID: " + strconv.Itoa(skeletonDataID))

	errEncode := json.NewEncoder(writer).Encode(newSkeletonData)
	if errEncode != nil {
		fmt.Println("error:", errEncode)
	}
}

func UpdateSkeletonData(writer http.ResponseWriter, request *http.Request) {
	skeletonDataID := getSidFromRequest(request)
	var updateSkeletonData dto.SkeletonData

	reqBody, errReadAll := io.ReadAll(request.Body)
	if errReadAll != nil {
		fmt.Println("error:", errReadAll)
	}

	errUnmarshal := json.Unmarshal(reqBody, &updateSkeletonData)
	if errUnmarshal != nil {
		fmt.Println("error:", errUnmarshal)
	}

	writer.WriteHeader(http.StatusOK)

	var newSkeletonData dto.SkeletonData
	newSkeletonData.SID = skeletonDataID
	newSkeletonData.Description = updateSkeletonData.Description
	log.Println("Update SkeletonData with ID: " + strconv.Itoa(skeletonDataID))

	errEncode := json.NewEncoder(writer).Encode(newSkeletonData)
	if errEncode != nil {
		fmt.Println("error:", errEncode)
	}
}

func DeleteSkeletonData(writer http.ResponseWriter, request *http.Request) {
	skeletonDataID := getSidFromRequest(request)
	writer.WriteHeader(http.StatusNoContent)

	log.Println("Delete SkeletonData with ID: " + strconv.Itoa(skeletonDataID))
}

func getSidFromRequest(req *http.Request) int {
	vars := mux.Vars(req)
	sid, _ := strconv.Atoi(vars["sid"])
	return sid
}
