package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rj-project/models"
)

var users = []models.Person{}

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func PostHandler(c *gin.Context) {
	var user models.Person

	name := c.Query("name")
	is_admin := c.Query("isAdmin")
	if is_admin != "" {
		is_admin = "true" //strconv.ParseBool(is_admin)
	} else {
		is_admin = "false"
	}

	user.Id = len(users) + 1
	user.Name = name

	users = append(users, user)

	c.JSON(200, user)
}

func DeleteHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, item := range users {
		if item.Id == id {
			users = append(users[:i], users[i+1:]...)
			return
		}
	}
	c.JSON(200, id)
}

func PutHandler(c *gin.Context) {
	var user models.Person
	id, _ := strconv.Atoi(c.Param("id"))

	for _, item := range users {
		if id == item.Id {

			user.Id = id
			name := c.Query("name")
			is_admin := c.Query("is_admin")

			if name != "" {
				user.Name = name
			} else {
				user.Name = item.Name
			}

			if is_admin != "" {
				user.IsAdmin, _ = strconv.ParseBool(is_admin)
			} else {
				user.IsAdmin = item.IsAdmin
			}

		}
	}

	c.JSON(200, user)
}

func GetHandler(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(500, gin.H{"ststus": "Server critical error"})
	}

	for _, item := range users {
		if id == item.Id {
			c.JSON(200, item)
			return
		}
	}

	c.JSON(404, models.Person{Id: 0, Name: "Anonymous", IsAdmin: false})
}
