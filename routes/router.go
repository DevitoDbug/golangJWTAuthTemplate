package routes

import (
	"fmt"
	"net/http"
)

func Router(w http.ResponseWriter, r *http.Request) {
	httpMethod := r.Method

	switch r.URL.Path {
	case "/register":
		if httpMethod == "POST" {
			// handle the registration
		} else {
			fmt.Fprint(w, "Invalid route", http.StatusBadRequest)
		}
		break
	case "/login":
		if httpMethod == "POST" {
			// handle the registration
		} else {
			fmt.Fprint(w, "Invalid route", http.StatusBadRequest)
		}
		break
	case "/protected":
		if httpMethod == "GET" {
			// handle the registration
		} else {
			fmt.Fprint(w, "Invalid route", http.StatusBadRequest)
		}
		break
	default:
		fmt.Fprint(w, "Invalid route", http.StatusBadRequest)
	}
}
