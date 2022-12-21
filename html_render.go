package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//  HTML渲染 应该一般用不上
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	//  router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	r.GET("/ginDemo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
