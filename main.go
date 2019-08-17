package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	utils "./utils"
	app "./app"
	controllers "./controllers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":" + port, router)
	utils.ParseError(err)
}