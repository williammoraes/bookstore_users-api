package app

import (
	"github.com/williammoraes/bookstore_users-api/controllers/ping"
	"github.com/williammoraes/bookstore_users-api/controllers/users"
)

func mapsUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
