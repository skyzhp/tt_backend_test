package handle

import (
	"backend_test/db_models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsers(c *gin.Context) {
	users := db_models.GetAllUser()
	fmt.Printf("users: %v \n", users)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    users,
	})
}

type CreateUserParam struct {
	Name string `json:"name"`
}

func CreateUser(c *gin.Context) {
	param := CreateUserParam{}
	c.BindJSON(&param)
	if param.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "wrong param!"})
		return
	}
	user := db_models.User{}
	user.Name = param.Name
	user.Type = db_models.UserType
	_, err := db_models.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "server error"})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}