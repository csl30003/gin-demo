package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//  使用JSONP向不同域的服务器请求数据 如果查询参数存在回调 则将回调添加到相应体中
	r := gin.Default()

	r.GET("/ginDemo", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		c.JSONP(http.StatusOK, data)
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
