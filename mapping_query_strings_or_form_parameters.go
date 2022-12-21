package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/ginDemo", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("%v %v", ids, names)
		c.String(http.StatusOK, "success")
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
