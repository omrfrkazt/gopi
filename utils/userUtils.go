package utils

import (
	"errors"
	. "testApi/models"
)

func GetActiveUsers(users []User) (err error,activeUsers []User) {
	if IsEmptyArray(users) {
		for _, user := range users {
			if user.IsActive {
				activeUsers = append(activeUsers, user)
			}
		}
	}
	err = errors.New("No active users found")
	return
}
