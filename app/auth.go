package app

import (
	"../models"
	"../utils"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		nonAuthEndpoints := []string{"/api/user/new", "/api/user/login"}
		requestPath := r.URL.Path
		for _, endpoint := range nonAuthEndpoints {
			if endpoint == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			response = utils.Message(false, "Auth token is missing.")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, response)
			return
		}

		tk := &models.Token{}

		//parse the received token using the claims from the models.Token{} struct
		token, err := jwt.ParseWithClaims(tokenHeader, tk, func(token *jwt.Token) (interface{}, error){
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			response = utils.Message(false, "Malformed auth token.")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, response)
			return
		}

		if !token.Valid {
			response = utils.Message(false, "Token is invalid.")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, response)
			return
		}

		fmt.Printf("User %", tk.UserId)
		fmt.Println(tk.UserId)
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	});
}