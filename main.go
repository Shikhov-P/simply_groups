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

	//staticFileDir := http.Dir("./assets/")
	//staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDir))
	//router.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/me/groups", controllers.GetCreatedGroupsFor).Methods("GET")
	router.HandleFunc("/api/groups/new", controllers.CreateGroup).Methods("POST")
	router.HandleFunc("/api/groups/delete/{groupId}", controllers.DeleteGroup).Methods("DELETE")
	router.HandleFunc("/api/groups/update/{groupId}", controllers.UpdateGroup).Methods("POST")

	router.Use(app.JwtAuthentication)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":" + port, router)
	utils.ParseError(err)
}