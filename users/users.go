package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func GetUsers(c *gin.Context) {
  users, err := GetUsersS()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
			"data":  nil,
		})
		return
	}

  c.JSON(http.StatusOK, gin.H{
    "error": nil,
    "data": users,
  })
}

func GetUser(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		// Handle the error if the conversion fails
		c.JSON(400, gin.H{
      "error": "Invalid ID format",
      "data": nil,
    })
		return
	}

  user, err := GetUserS(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
      "error": err,
      "data": nil,
    })
		return
	}

	c.JSON(http.StatusOK, gin.H{
    "error": nil,
		"data": user,
	})
}
