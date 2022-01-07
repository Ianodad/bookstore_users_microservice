package app

import (
	"github.com/ianodad/bookstore_users_microservice/controllers/ping"
	"github.com/ianodad/bookstore_users_microservice/controllers/users"
)	
func mapUrls(){
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users/", users.CreateUser)
}

