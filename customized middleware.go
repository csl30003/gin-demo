package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		//  可以设置变量
		c.Set("key", "value")

		//  请求前

		c.Next()

		//  请求后

		latency := time.Since(t)
		log.Println(latency)

		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.New()

	r.Use(logger())

	r.GET("/ginDemo", func(c *gin.Context) {
		val := c.MustGet("key").(string)
		log.Println(val)

		c.String(http.StatusOK, "ok")
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
