package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/ginDemo", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	log.Fatal(autotls.Run(r, "example1.com", "example2.com"))
}
