package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/longAsync", func(c *gin.Context) {
		//  用副本
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println(cCp.Request.URL.Path)
		}()
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
