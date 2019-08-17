package app

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
	utils "../utils"
	models "../models"
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

		splittedTokenHeader := strings.Split(tokenHeader, " ")
		if len(splittedTokenHeader) != 2 {
			response = utils.Message(false, "Malformed auth token.")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, response)
			return
		}

		tokenPart := splittedTokenHeader[1]
		tk := &models.Token{}

		//parse the received token using the claims from the models.Token{} struct
		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error){
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

		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	});
}