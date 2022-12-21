package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func startPage(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		c.String(http.StatusOK, "success")
	} else {
		c.String(http.StatusOK, "failed")
	}
}

func main() {
	r := gin.Default()
	r.GET("/testing", startPage)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
