package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DevitoDbug/golangJWTAuthTemplate/models"
	"github.com/DevitoDbug/golangJWTAuthTemplate/repository"
	"github.com/DevitoDbug/golangJWTAuthTemplate/utils"
	"github.com/go-playground/validator/v10"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		customError := utils.Error{
			Context: "Register@authController.go",
			Info:    "Failed to decode request body",
		}

		http.Error(w, customError.Error(), http.StatusInternalServerError)
	}

	// struct validation
	validationError := utils.Validate.Struct(user)
	if validationError != nil {
		var errorResponse []utils.ValidationError

		if validationError, ok := validationError.(validator.ValidationErrors); ok {
			for _, err := range validationError {
				errorValue := utils.ValidationError{
					Field:   err.Field(),
					Tag:     err.Tag(),
					Message: err.Error(),
				}
				errorResponse = append(errorResponse, errorValue)
			}
		} else {
			http.Error(w, "Unexpected validation error", http.StatusInternalServerError)
			return
		}

		http.Error(w, "Validation failed", http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(errorResponse)
		if err != nil {
			http.Error(w, "Failed to write response payload", http.StatusInternalServerError)
			return
		}
	}

	hashedPassword, err := utils.GenerateHash(user.Password)
	if err != nil {
		http.Error(w, "Failed to generate password hash", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	repository.Storage[user.Email] = user
	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(user.GetPublicUserInfo())
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

type loginCredentials struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var data loginCredentials

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Printf("Error while decoding the request: \nError: %v", err)
		http.Error(w, "Failed to decode response body", http.StatusBadRequest)
		return
	}

	// struct validation
	validationError := utils.Validate.Struct(data)
	if validationError != nil {
		var errorResponse []utils.ValidationError

		if validationError, ok := validationError.(validator.ValidationErrors); ok {
			for _, err := range validationError {
				errorValue := utils.ValidationError{
					Field:   err.Field(),
					Tag:     err.Tag(),
					Message: err.Error(),
				}
				errorResponse = append(errorResponse, errorValue)
			}
		} else {
			fmt.Printf("Error while doing validation: %v", validationError)
			http.Error(w, "Unexpected validation error", http.StatusInternalServerError)
			return
		}

		http.Error(w, "Validation failed", http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(errorResponse)
		if err != nil {
			http.Error(w, "Failed to decode response body", http.StatusInternalServerError)
			return
		}
	}

	user, ok := repository.Storage[data.Email]
	passwordIsValid := utils.IsPasswordValid(user.Password, data.Password)
	if !ok || !passwordIsValid {
		http.Error(w, "Invalid user name or password", http.StatusUnauthorized)
		return
	}

	jwtToken, err := utils.CreateToken(user.Email)
	if err != nil {
		fmt.Printf("Error while creating token.\nError: %v", err)
		http.Error(w, "Failed to create auth token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt_token",
		Value:    jwtToken,
		MaxAge:   3600,
		Domain:   "localhost",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
	})

	w.Header().Set("content-type", "application/json")
	w.Write([]byte("Login successful"))
}

func LogOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt_token",
		Value:    "",
		MaxAge:   -1,
		Domain:   "localhost",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
	})

	w.Header().Set("content-type", "application/json")
	w.Write([]byte("Log out successful"))
}
