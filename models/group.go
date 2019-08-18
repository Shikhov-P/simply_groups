package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	utils "../utils"
)

type Group struct {
	gorm.Model
	Name string `json:"name"`
	Description string `json:"description"`
	UserId uint `json:"user_id"`
}

//TODO: how to check user existence
func (group *Group) Validate() (map[string]interface{}, bool){
	if !(utils.IsAlphaNumeric(group.Name) && utils.IsAlphaNumeric(group.Description)){
		return utils.Message(false, "Group name and description must contain letters only."), false
	}

	if group.UserId <= 0 {
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
	if err == nil {
		fmt.Println(err)
		return nil
	}
	return groups
}
