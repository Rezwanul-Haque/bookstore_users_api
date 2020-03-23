package users

import (
	"encoding/json"
	"github.com/rezwanul-haque/microservices_in_golang/bookstore_users_api/utils/helpers"
)

type PublicUser struct {
	Id       int64  `json:"user_id"`
	FullName string `json:"first_name"`
	Status   string `json:"status"`
}

type PrivateUser struct {
	Id          int64  `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	// Password    string `json:"password"`
}

func (users Users) Marshall(isPublic bool) interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:       user.Id,
			FullName: helpers.FullName(user.FirstName, user.LastName),
			Status:   user.Status,
		}
	}
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}
