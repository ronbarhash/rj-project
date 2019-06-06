package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rj-project/models"
)

var Users = []models.Person{}

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Users)
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

	user.Id = len(Users) + 1
	user.Name = name

	Users = append(Users, user)

	c.JSON(200, user)
}

func DeleteHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, item := range Users {
		if item.Id == id {
			Users = append(Users[:i], Users[i+1:]...)
			return
		}
	}
	c.JSON(200, id)
}

func PutHandler(c *gin.Context) {
	var user models.Person

	id, _ := strconv.Atoi(c.Param("id"))

	for i, item := range Users {
		if id == item.Id {

			users := Users
			user.Id = id
			name := c.Query("name")
			is_admin := c.Query("is_admin")

			if name != "" {
				user.Name = name
			}

			if is_admin != "" {
				user.IsAdmin, _ = strconv.ParseBool(is_admin)
			}

			Users = append(Users[:i], user)
			Users = append(Users, users[i+1:]...)

		}
	}

	c.JSON(200, id)
}

func GetHandler(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(500, gin.H{"ststus": "Server critical error"})
	}

	for _, item := range Users {
		if id == item.Id {
			c.JSON(200, item)
			return
		}
	}

	c.JSON(404, models.Person{Id: 0, Name: "Anonymous", IsAdmin: false})
}
