package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/test1", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	})

	r.POST("/test2", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/foo")
	})

	r.GET("/test3", func(c *gin.Context) {
		//  路由重定向
		c.Request.URL.Path = "/test4"
		r.HandleContext(c)
	})
	r.GET("/test4", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello")
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
