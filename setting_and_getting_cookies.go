package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "not set"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		log.Println(cookie)
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
