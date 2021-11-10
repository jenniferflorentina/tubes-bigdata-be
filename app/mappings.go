package app

import (
	"github.com/tubes-bigdata/controllers/users"
)

func MapUrls() {
	router.GET("/users", users.GetAllUser)
	router.GET("/users/:id", users.FindUser)
	router.DELETE("/users/:id", users.DeleteUser)
	router.PUT("/users/:id", users.UpdateUser)
	router.POST("/users", users.CreateUser)
}
