package controllers

import (
	"../models"
	"../utils"
	"encoding/json"
	//"github.com/gorilla/mux"
	"net/http"
	//"strconv"
)

var CreateGroup = func(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user") . (uint)
	group := &models.Group{}
	err := json.NewDecoder(r.Body).Decode(group)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request."))
		return
	}
	group.UserId = user
	response := group.Create()
	utils.Respond(w, response)
}


var GetCreatedGroupsFor = func(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(uint)
	data := models.GetCreatedGroups(userId)
	response := utils.Message(true, "Success.")
	response["data"] = data
	utils.Respond(w, response)
}
