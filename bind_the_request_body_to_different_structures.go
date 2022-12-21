package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func main() {
	r := gin.Default()

	r.POST("ginDemo", func(c *gin.Context) {
		objA := formA{}
		objB := formB{}

		if errA := c.ShouldBind(&objA); errA == nil {
			c.String(http.StatusOK, "the body should be formA")
		} else if errB := c.ShouldBind(&objB); errB == nil {
			c.String(http.StatusOK, "the body should be formB")
		} else {
			c.String(http.StatusOK, "should bind err")
		}
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
