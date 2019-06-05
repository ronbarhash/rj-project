package main

import (
	"net/http"

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
	r.PUT("/{id}", putHandler)
	r.DELETE("/{id}", deleteHandler)

	r.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}
func postHandler(c *gin.Context) {
}
func putHandler(c *gin.Context)    {}
func deleteHandler(c *gin.Context) {}
