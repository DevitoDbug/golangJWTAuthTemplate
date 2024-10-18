package routes

import (
	"fmt"
	"net/http"

	"github.com/DevitoDbug/golangJWTAuthTemplate/controllers"
	"github.com/DevitoDbug/golangJWTAuthTemplate/middleware"
)

func Router(w http.ResponseWriter, r *http.Request) {
	httpMethod := r.Method

	switch r.URL.Path {
	case "/register":
		if httpMethod == "POST" {
			controllers.Register(w, r)
		} else {
			fmt.Fprint(w, "Invalid route", http.StatusBadRequest)
		}
		break
	case "/login":
		if httpMethod == "POST" {
			controllers.Login(w, r)
		} else {
			fmt.Fprint(w, "Invalid route", http.StatusBadRequest)
		}
		break
	case "/logout":
		if httpMethod == "POST" {
			controllers.LogOut(w, r)
		} else {
			fmt.Fprint(w, "Invalid route", http.StatusBadRequest)
		}
		break
	case "/get-all":
		if httpMethod == "GET" {
			middleware.Auth(http.HandlerFunc(controllers.GetAllUsers)).ServeHTTP(w, r)
		} else {
			fmt.Fprint(w, "Invalid route", http.StatusBadRequest)
		}
		break
	default:
		fmt.Fprint(w, "Invalid route", http.StatusBadRequest)
	}
}
