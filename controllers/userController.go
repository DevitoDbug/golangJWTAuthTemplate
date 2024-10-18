package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/DevitoDbug/golangJWTAuthTemplate/models"
	"github.com/DevitoDbug/golangJWTAuthTemplate/repository"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	_ = r
	var responseData []models.PublicUserInfo
	allStudents := repository.Storage
	for _, student := range allStudents {
		responseData = append(responseData, student.GetPublicUserInfo())
	}

	w.Header().Set("Content-type", "application/json")
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
