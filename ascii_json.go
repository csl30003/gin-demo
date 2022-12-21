package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//  使用AsciiJSON生成具有转义的非ASCII字符的ASCII-only JSON
	r := gin.Default()

	r.GET("/ginDemo", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO",
			"tag":  "<br>",
		}

		c.AsciiJSON(http.StatusOK, data)
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
