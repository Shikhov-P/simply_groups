package controllers

import (
	"encoding/json"
	"net/http"
	models "../models"
	utils "../utils"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.User{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request."))
		return
	}
	response := account.Create()
	utils.Respond(w, response)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	account := &models.User{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request."))
		return
	}
	response := models.Login(account.Email, account.Password)
	utils.Respond(w, response)
}