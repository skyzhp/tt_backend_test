package main

import (
	"backend_test/handle"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/users", handle.GetAllUsers)
	r.POST("/users", handle.CreateUser)
	r.PUT("/users/:user_id/relationships/:other_user_id", handle.UpdateRelationships)
	r.GET("/users/:user_id/relationships", handle.GetRelationships)
	r.Run(":80")
}
