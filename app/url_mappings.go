package app

import (
	"github.com/rezwanul-haque/microservices_in_golang/bookstore_users_api/controllers/ping"
	"github.com/rezwanul-haque/microservices_in_golang/bookstore_users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	// GET /internal/users/search?status=active
	router.GET("/internal/users/search", users.Search)
}
