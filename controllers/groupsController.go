package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	models "../models"
	utils "../utils"
	"strconv"
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
	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Group id must be an integer."))
		return
	}
	data := models.GetCreatedGroups(uint(userId))
	response := utils.Message(true, "Success.")
	response["data"] = data
	utils.Respond(w, response)
}
