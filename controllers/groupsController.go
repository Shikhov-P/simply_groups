package controllers

import (
	"../models"
	"../utils"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
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
	group.UserID = user
	response := group.Create()
	utils.Respond(w, response)
}

//TODO: check if url parameter is not a number
var DeleteGroup = func(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user") . (uint)
	groupId, _ := strconv.ParseUint(mux.Vars(r)["groupId"], 10, 32)
	response := models.Delete(userId, uint(groupId))
	utils.Respond(w, response)
}

var UpdateGroup = func(w http.ResponseWriter, r *http.Request) {
	var group models.Group
	user := r.Context().Value("user") . (uint)
	err := json.NewDecoder(r.Body).Decode(&group)
	group.UserID = user
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request."))
		return
	}
	groupId, _ := strconv.ParseUint(mux.Vars(r)["groupId"], 10, 32)
	response := models.Update(uint(groupId), user, group)
	utils.Respond(w, response)
}

var GetCreatedGroupsFor = func(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("user").(uint)
	data := models.GetCreatedGroups(userId)
	response := utils.Message(true, "Success.")
	response["data"] = data
	utils.Respond(w, response)
}
