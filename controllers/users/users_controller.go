package users

import (
	"net/http"

	"github.com/tubes-bigdata/domain"
	"github.com/tubes-bigdata/services"
	"github.com/tubes-bigdata/utils"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser domain.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		restErr := utils.BadRequest("Invalid json.")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, restErr := services.CreateUser(&newUser)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusCreated, user)
}

func GetAllUser(c *gin.Context) {
	user, restErr := services.GetAll()
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func FindUser(c *gin.Context) {
	id := c.Param("id")
	user, restErr := services.FindUser(id)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	restErr := services.DeleteUser(id)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"isRemoved": true})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var newUser domain.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		restErr := utils.BadRequest("Invalid json.")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, restErr := services.UpdateUser(id, &newUser)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
