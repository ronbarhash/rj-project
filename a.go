package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	isAdmin bool   `json:"is_admin"`
}

var users []Person

func main() {

	users = []Person{
		Person{Id: 1, Name: "John Smith", isAdmin: false},
		Person{Id: 2, Name: "John Dou", isAdmin: false},
		Person{Id: 3, Name: "Angela Smith", isAdmin: true},
		Person{Id: 4, Name: "Ron Smith", isAdmin: true},
	}

	r := gin.Default()
	r.GET("/", rootHandler)
	r.POST("/post", postHandler)
	r.DELETE("/:id", deleteHandler)
	r.PUT("/:id", putHandler)
	r.GET("/:id", getHandler)

	r.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func postHandler(c *gin.Context) {
	var user Person

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

func deleteHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, item := range users {
		if item.Id == id {
			users = append(users[:i], users[i+1]...)
			return
		}
	}
	c.JSON(200, item)
}

func putHandler(c *gin.Context) {
	var user Person

	c.JSON(200, user)
}

func getHandler(c *gin.Context) {

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

	c.JSON(404, Person{Id: 0, Name: "Anonymous", isAdmin: false})
}
