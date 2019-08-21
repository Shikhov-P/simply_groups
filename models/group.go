package models

import (
	"../utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type Group struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint   `json:"account_id"`
}

//TODO: how to check user existence
//TODO: injection prevention
func (group *Group) Validate() (map[string]interface{}, bool){
	if !(utils.IsAlphaNumeric(group.Name) && utils.IsAlphaNumeric(group.Description)){
		return utils.Message(false, "Group name and description must contain letters only."), false
	}

	if strings.TrimSpace(group.Name) == "" && strings.TrimSpace(group.Description) == "" {
		return utils.Message(false, "Group name and description must be populated."), false
	}

	if group.UserID <= 0 {
		return utils.Message(false, "Try another user."), false
	}

	return utils.Message(true, "Group is validated."), true
}



func (group *Group) Create() map[string]interface{} {
	if response, isGroupValid := group.Validate(); !isGroupValid {
		return response
	}

	GetDB().Create(group)

	response := utils.Message(true, "Group is created.")
	response["group"] = group
	return response
}

func GetCreatedGroups(userId uint) []*Group {
	groups := make([]*Group, 0)
	err := GetDB().Table("groups").Where("user_id = ?", userId).Find(&groups).Error
	if err != nil {
		fmt.Println("Error", err)
		return nil
	}
	return groups
}

func Delete(userId, groupId uint) map[string]interface{} {
	var response map[string]interface{}
	if groupId == 0 {
		response = utils.Message(true, "Delete operation failed.")
		return response
	}
	rowsAffected := GetDB().Delete(Group{}, "user_id = ? AND id = ?", userId, groupId).RowsAffected
	if rowsAffected == 0 {
		response = utils.Message(true, "Delete operation failed.")
	} else {
		response = utils.Message(true, "Group is deleted.")
	}
	return response
}

func IsUserGroupOwner(groupId, userId uint) bool {
	if groupId == 0 {
		return false
	}
	var group Group
	GetDB().Where(Group{Model: gorm.Model{ID: groupId}}).First(&group)
	if group.UserID == userId {
		return true
	} else {
		return false
	}
}

func Update(groupId, userId uint, group Group) map[string]interface{} {
	var response map[string]interface{}
	if !IsUserGroupOwner(groupId, userId) {
		response = utils.Message(true, "Unauthorized user.")
		return response
	}
	if response, isGroupValid := group.Validate(); !isGroupValid {
		return response
	}
	GetDB().Model(&Group{}).Where("id = ?", groupId).Updates(group)
	response = utils.Message(true, "Group is updated.")
	return response
}
