package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	r.POST("/ginDemo", func(c *gin.Context) {
		var form LoginForm

		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(http.StatusOK, gin.H{
					"status": "login success",
				})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"status": "login failed",
				})
			}
		}
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
