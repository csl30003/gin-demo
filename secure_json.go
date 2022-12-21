package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/ginDemo", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		c.SecureJSON(http.StatusOK, names)
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
