package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/ginDemo", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"id":      id,
			"page":    page,
			"name":    name,
			"message": message,
		})
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
