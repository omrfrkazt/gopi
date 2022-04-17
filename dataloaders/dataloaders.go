package dataloaders

import (
	"encoding/json"

	. "testApi/models"
	util "testApi/utils"
)

func LoadUsers(onlyActive bool) []User {
	bytes, _ := util.ReadFile("../json/users.json")
	var users []User
	json.Unmarshal([]byte(bytes), &users)
	if onlyActive {
		err, data := util.GetActiveUsers(users)
		if err != nil {
			util.CheckError(err)
		} else {
			return data
		}
	}
	return users
}
