package app

import (
	"github.com/rezwanul-haque/microservices_in_golang/bookstore_users_api/controllers/ping"
	"github.com/rezwanul-haque/microservices_in_golang/bookstore_users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
