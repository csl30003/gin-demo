package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/ginDemo", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "csl")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
