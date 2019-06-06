package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rj-project/models"
)
import rr "github.com/rj-project/router"

var users []models.Person

func main() {

	users = []models.Person{
		models.Person{Id: 1, Name: "John Smith", IsAdmin: false},
		models.Person{Id: 2, Name: "John Dou", IsAdmin: false},
		models.Person{Id: 3, Name: "Angela Smith", IsAdmin: true},
		models.Person{Id: 4, Name: "Ron Smith", IsAdmin: true},
	}

	r := gin.Default()
	r.GET("/", rr.RootHandler)
	r.POST("/post", rr.PostHandler)
	r.DELETE("/:id", rr.DeleteHandler)
	r.PUT("/:id", rr.PutHandler)
	r.GET("/:id", rr.GetHandler)

	r.Run()
}
