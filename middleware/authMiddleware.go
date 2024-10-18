package middleware

import (
	"fmt"
	"net/http"

	"github.com/DevitoDbug/golangJWTAuthTemplate/utils"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt_token")
		if err != nil {
			http.Error(w, "Token missing in cookie", http.StatusBadRequest)
			return
		}

		token, err := utils.VerifyToken(cookie.Value)
		if err != nil {
			http.Error(w, "Token verification failed", http.StatusBadRequest)
			return
		}

		fmt.Printf("Token verified successfully. \nClaims:%v", token.Claims)

		next.ServeHTTP(w, r)
	})
}
